package inbox

import (
	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectInitAchievement        = "inbox.achievement.init"
	SubjectRecalculateAchievement = "inbox.achievement.recalculate"
)

type AchievementType string

const (
	AchievementTypeAppInfo      AchievementType = "app_info"
	AchievementTypeVoteVerified AchievementType = "vote_verified"
)

type AchievementInitEvent struct {
	UserID uuid.UUID `json:"user_id"`
}

type AchievementInitHandler = events.Handler[AchievementInitEvent]

type AchievementRecalculateEvent struct {
	UserID uuid.UUID       `json:"user_id"`
	Type   AchievementType `json:"type"`
}

type AchievementRecalculateHandler = events.Handler[AchievementRecalculateEvent]
