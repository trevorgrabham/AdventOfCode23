### Linked List?

- Use `Sprinklers` to keep track of the sprinkler data for the problem

  - `Sprinklers.Groups` will be the head to the linked list of `Group` nodes
  - `Sprinklers.NumGroups` will be used to keep track of the length of the linked list
  - `Sprinklers.ReferenceNumbers` will be an array of the reference number and order of the contiguous broken sprinklers from the problem

- Use `Group` to represent a contiguous group of `Sprinkler`s in identical `Sprinkler.State`. It will be used as a `Node` for the `Sprinklers` linked list

  - `Group.Type` will be of the same type as `Sprinkler.State`, and will be used to keep track of the `Sprinkler.State` value for all of the `Group`
  - `Group.NumMembers` to keep track of how long the contiguous group is
    - If the `Group.Type == Working` then the length of the `Group` doesn't really matter
    - If the `Group.Type == Broken || Group.Type == Unknown` then the length will be important
  - `Group.left`, `Group.right` to keep track of their siblings
    - These values will be used in a `Group` function that will attempt to borrow a `Sprinkler` with `Sprinkler.State == Broken` from a neighbour -`Group.NumSolutions` will be

- `func (s Sprinklers)

-`func (g *Group) LeftTerminated() bool` will only be used when `g.Type == Broken` and will check if the

- `func (g *Group) BorrowLeft(target int) bool` will be used to attempt to borrow from the group to the left

  - If `g.left == nil` then we obviously can't borrow from them
  - If `g.Type != Broken` then we have no reason to borrow from any neighbour
  - If left `g.left` is `left.Type == Working` then we cannot borrow from them
  - If `g.left` is `left.Type == Broken` then we must make sure that `g.NumMembers + left.NumMembers` isn't greater than `target`
  - Otherwise `g.left` is `left.Type == Unknown`
    - If `g.left` has just a single member
      - Set `g.left` to its old neighbours neighbour and set the new left neighbours right to `g`
    - Otherwise remove a neighbour from `g.left` and add it to `g`

- `func (g *Group) BorrowRight() bool` exactly the same as `BorrowLeft()`
