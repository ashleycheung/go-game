package physics

const DefaultSplitAmount = 5
const DefaultMaxDepth = 20

// Creates a quad tree given a slice of bodies
func NewQuadTreeFromBodies(bodies []*Body, splitAmount, maxDepth int) *QuadTree {
	region := BBox{}
	for _, b := range bodies {
		var bodyTopLeft Vector
		var bodyBottomRight Vector

		switch b.Shape.GetType() {
		case CircleType:
			circle := b.Shape.(Circle)
			bodyTopLeft = b.Position.Subtract(Vector{
				X: circle.Radius,
				Y: circle.Radius,
			})
			bodyBottomRight = b.Position.Add(Vector{
				X: circle.Radius,
				Y: circle.Radius,
			})
		case RectangleType:
			rect := b.Shape.(Rectangle)
			bodyTopLeft, bodyBottomRight = RectangleCorners(b.Position, rect)
		default:
			panic("unsupported type " + b.Shape.GetType())
		}
		// Update region if necessary
		if bodyTopLeft.X < region.TopLeft.X {
			region.TopLeft.X = bodyTopLeft.X
		}
		if bodyTopLeft.Y < region.TopLeft.Y {
			region.TopLeft.Y = bodyTopLeft.Y
		}
		if bodyBottomRight.X > region.BottomRight.X {
			region.BottomRight.X = bodyBottomRight.X
		}
		if bodyBottomRight.Y > region.BottomRight.Y {
			region.BottomRight.Y = bodyBottomRight.Y
		}
	}

	// Build quadtree
	qTree := NewQuadTree(region, splitAmount, maxDepth)
	for _, b := range bodies {
		qTree.AddBody(b)
	}
	return qTree
}

// Creates a new quad tree
func NewQuadTree(region BBox, splitAmount int, maxDepth int) *QuadTree {
	if splitAmount <= 1 {
		panic("split amount cant be 1 or less as it will lead to infinite loop")
	}
	tree := &QuadTree{
		splitAmount:       splitAmount,
		maxDepth:          maxDepth,
		bodyQuadTreeNodes: map[*Body]map[*QuadTreeNode]bool{},
	}
	tree.rootNode = NewQuadTreeNode(region, tree)
	return tree
}

// Implements a quad tree
type QuadTree struct {
	// The amount of bodies
	// contained before splitting
	// the tree
	splitAmount int

	// Root node
	rootNode *QuadTreeNode

	// Maps the body to the
	// nodes it is a part of
	bodyQuadTreeNodes map[*Body]map[*QuadTreeNode]bool

	// The max depth before no longer splitting
	maxDepth int
}

// Gets all the bboxes in the quadtree
func (qTree *QuadTree) GetBBoxes() []BBox {
	bboxes := []BBox{}
	var dfs func(node *QuadTreeNode)
	dfs = func(node *QuadTreeNode) {
		if !node.hasSplit {
			bboxes = append(bboxes, node.Region)
		} else {
			dfs(node.topLeft)
			dfs(node.topRight)
			dfs(node.bottomLeft)
			dfs(node.bottomRight)
		}
	}
	dfs(qTree.rootNode)
	return bboxes
}

// Adds abody into the tree
func (qTree *QuadTree) AddBody(body *Body) {
	qTree.rootNode.AddBody(body)
}

// Get all the bodies that the given
// body are close to
func (qTree *QuadTree) GetNeighbours(body *Body) []*Body {
	nodeSet, exists := qTree.bodyQuadTreeNodes[body]
	if exists {
		neighbours := map[*Body]bool{}
		for node := range nodeSet {
			for neighBody := range node.bodies {
				if neighBody != body {
					neighbours[neighBody] = true
				}
			}
		}
		neighBodies := []*Body{}
		for nBody := range neighbours {
			neighBodies = append(neighBodies, nBody)
		}
		return neighBodies
	}
	return []*Body{}
}

// Used by the quad tree node to cache
// the body to a given node
func (qTree *QuadTree) cacheBodyToNode(body *Body, node *QuadTreeNode) {
	nodeSet, exists := qTree.bodyQuadTreeNodes[body]
	if exists {
		nodeSet[node] = true
	} else {
		qTree.bodyQuadTreeNodes[body] = map[*QuadTreeNode]bool{
			node: true,
		}
	}
}

// Uncache a body to node
func (qTree *QuadTree) removeCacheBodyToNode(body *Body, node *QuadTreeNode) {
	nodeSet, exists := qTree.bodyQuadTreeNodes[body]
	if exists {
		delete(nodeSet, node)
	}
}

// Creates a new quad tree node
func NewQuadTreeNode(region BBox, tree *QuadTree) *QuadTreeNode {
	return &QuadTreeNode{
		bodies: map[*Body]bool{},
		tree:   tree,
		Region: region,
	}
}

// A node in a quadtree
type QuadTreeNode struct {
	// All bodies stored
	// in this body
	bodies map[*Body]bool

	// The tree itself
	tree *QuadTree

	// The region that this
	// quad tree covers
	Region BBox

	// Whether the quad tree
	// has split yet
	hasSplit bool

	// The parent quad tree
	parent *QuadTreeNode

	// Current depth of the tree
	depth int

	// The children quadtrees
	topLeft     *QuadTreeNode
	topRight    *QuadTreeNode
	bottomLeft  *QuadTreeNode
	bottomRight *QuadTreeNode
}

// Adds a body
func (qNode *QuadTreeNode) AddBody(b *Body) {
	// Check if the body is within this region.
	// If not return
	switch b.Shape.GetType() {
	case RectangleType:
		rect := b.Shape.(Rectangle)
		regionSize, regionPos := qNode.Region.ToSizePosition()
		// Not inside region so return
		if !RectangleRectangleCollision(rect.Size, b.Position, regionSize, regionPos) {
			return
		}
	case CircleType:
		circle := b.Shape.(Circle)
		regionSize, regionPos := qNode.Region.ToSizePosition()
		// Not inside region so return
		if !CircleRectangleCollision(circle.Radius, b.Position, regionSize, regionPos) {
			return
		}
	default:
		panic("unsupported type: " + b.Shape.GetType())
	}

	// Split if necessary and less than depth
	if len(qNode.bodies)+1 >= qNode.tree.splitAmount && qNode.depth+1 < qNode.tree.maxDepth {
		qNode.split()
	}

	// If already split
	// give it to the children
	if qNode.hasSplit {
		qNode.topLeft.AddBody(b)
		qNode.topRight.AddBody(b)
		qNode.bottomLeft.AddBody(b)
		qNode.bottomRight.AddBody(b)
	} else {
		// Not split so add to self
		qNode.bodies[b] = true
		// Cache to body
		qNode.tree.cacheBodyToNode(b, qNode)
	}
}

// Splits the quad tree.
// If already split, does nothing
func (qNode *QuadTreeNode) split() {
	// If already split
	// do nothing
	if qNode.hasSplit {
		return
	}

	// Split
	qNode.hasSplit = true

	// Create subtrees

	// Top Left
	{
		topLeftBBox := BBox{}
		topLeftBBox.TopLeft = qNode.Region.TopLeft
		topLeftBBox.BottomRight = topLeftBBox.TopLeft.Add(qNode.Region.Size().Scale(0.5))
		qNode.topLeft = NewQuadTreeNode(topLeftBBox, qNode.tree)
		qNode.topLeft.depth = qNode.depth + 1
		qNode.topLeft.parent = qNode
	}

	// Top Right
	{
		topRightBBox := BBox{}
		topRightBBox.TopLeft = qNode.Region.TopLeft.Add(Vector{X: qNode.Region.Size().Scale(0.5).X})
		topRightBBox.BottomRight = topRightBBox.TopLeft.Add(qNode.Region.Size().Scale(0.5))
		qNode.topRight = NewQuadTreeNode(topRightBBox, qNode.tree)
		qNode.topRight.depth = qNode.depth + 1
		qNode.topRight.parent = qNode
	}

	// Bottom Left
	{
		bottomLeftBBox := BBox{}
		bottomLeftBBox.TopLeft = qNode.Region.TopLeft.Add(Vector{Y: qNode.Region.Size().Scale(0.5).Y})
		bottomLeftBBox.BottomRight = bottomLeftBBox.TopLeft.Add(qNode.Region.Size().Scale(0.5))
		qNode.bottomLeft = NewQuadTreeNode(bottomLeftBBox, qNode.tree)
		qNode.bottomLeft.depth = qNode.depth + 1
		qNode.bottomLeft.parent = qNode
	}

	// Bottom Right
	{
		bottomRightBBox := BBox{}
		bottomRightBBox.TopLeft = qNode.Region.TopLeft.Add(qNode.Region.Size().Scale(0.5))
		bottomRightBBox.BottomRight = qNode.Region.BottomRight
		qNode.bottomRight = NewQuadTreeNode(bottomRightBBox, qNode.tree)
		qNode.bottomRight.depth = qNode.depth + 1
		qNode.bottomRight.parent = qNode
	}

	// For each body in current tree, add it to child
	// tree
	for b := range qNode.bodies {
		// Remove current body from
		// current node cache
		qNode.tree.removeCacheBodyToNode(b, qNode)

		// Add to child body
		qNode.topLeft.AddBody(b)
		qNode.topRight.AddBody(b)
		qNode.bottomLeft.AddBody(b)
		qNode.bottomRight.AddBody(b)
	}

	// Clear current quadtree
	qNode.bodies = map[*Body]bool{}
}
