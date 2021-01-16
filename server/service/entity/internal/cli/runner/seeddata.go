package runner

import (
	"github.com/urfave/cli/v2"

	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/harness"
	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/record"
)

// LoadSeedData -
func (rnr *Runner) LoadSeedData(c *cli.Context) error {

	rnr.Log.Info("** Load Seed Data **")

	// harness
	config := harness.DataConfig{

		EntityConfig: []harness.EntityConfig{
			// Dark Armoured
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeMage,
					Name:         "Dark Armoured",
					Strength:     18,
					Dexterity:    10,
					Intelligence: 10,
					Coins:        50,
				},
			},
			// Druid
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeMage,
					Name:         "Druid",
					Strength:     14,
					Dexterity:    14,
					Intelligence: 10,
					Coins:        50,
				},
			},
			// Fairy
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeMage,
					Name:         "Fairy",
					Strength:     10,
					Dexterity:    14,
					Intelligence: 14,
					Coins:        50,
				},
			},
			// Necromancer
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeMage,
					Name:         "Necromancer",
					Strength:     14,
					Dexterity:    10,
					Intelligence: 14,
					Coins:        50,
				},
			},
		},
	}

	h, err := harness.NewTesting(config)
	if err != nil {
		rnr.Log.Warn("Failed new testing harness >%v<", err)
		return err
	}

	// harness commit data
	h.CommitData = true

	err = h.Setup()
	if err != nil {
		rnr.Log.Warn("Failed testing harness setup >%v<", err)
		return err
	}

	rnr.Log.Info("All done")

	return nil
}
