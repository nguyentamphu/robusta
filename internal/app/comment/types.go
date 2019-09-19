package comment

type (
	FindRequest struct {
		Offset      int      `json:"offset,omitempty"`
		Limit       int      `json:"limit,omitempty"`
		Target      string   `json:"target,omitempty"`
		ReplyToID   string   `json:"reply_to_id,omitempty"`
		ThreadID    string   `json:"thread_id,omitempty"`
		CreatedByID string   `json:"created_by_id,omitempty"`
		SortBy      []string `json:"sort_by,omitempty"`
	}
)

// Actions policy
const (
	ActionCreate = "comment:create"
	ActionDelete = "comment:delete"
	ActionUpdate = "comment:update"
	ActionRead   = "coment:read"
)

// Policy object name
const (
	PolicyObject = "comment"
)
