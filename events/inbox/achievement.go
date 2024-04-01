package inbox

import (
	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectRecalculateAchievement = "inbox.achievement.recalculate"
)

type AchievementType string

const (
	AchievementTypeAppInfo      AchievementType = "app_info"
	AchievementTypeVoteVerified AchievementType = "vote_verified"
)

type AchievementRecalculateEvent struct {
	UserID uuid.UUID       `json:"user_id"`
	Type   AchievementType `json:"type"`
}

type AchievementRecalculateHandler = events.Handler[AchievementRecalculateEvent]
