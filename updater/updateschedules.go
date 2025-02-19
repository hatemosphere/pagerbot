package updater

import (
	"sync"

	"github.com/hatemosphere/pagerbot/config"
	log "github.com/sirupsen/logrus"
)

// Updates the schedules from Pagerduty, check that all schedules listed
// in config exist
func (u *Updater) updateSchedules(w *sync.WaitGroup) {
	defer w.Done()
	pdSchedules, err := u.Pagerduty.Schedules()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warning("Could not fetch schedules from Pagerduty")
		return
	}

	var schds ScheduleList
	for i := range pdSchedules {
		schds.schedules = append(schds.schedules, &pdSchedules[i])
	}

	u.Schedules = &schds

	for _, group := range config.Config.Groups {
		for _, schdId := range group.Schedules {
			s := u.Schedules.ById(schdId)
			if s == nil || s.Id == "" {
				log.WithFields(log.Fields{
					"scheduleId": schdId,
					"group":      group.Name,
				}).Warning("Could not find schedule specified in group")
			}
		}
	}
}
