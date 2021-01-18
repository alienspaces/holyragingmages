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
			// Mage - Dark Armoured
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterMage,
					Name:         "Dark Armoured",
					Avatar:       record.MageAvatarDarkArmoured,
					Strength:     16,
					Dexterity:    12,
					Intelligence: 10,
					Coins:        50,
				},
			},
			// Mage - Red Stripe Druid
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterMage,
					Name:         "Druid",
					Avatar:       record.MageAvatarRedStripeDruid,
					Strength:     14,
					Dexterity:    14,
					Intelligence: 10,
					Coins:        50,
				},
			},
			// Mage - Red Fairy
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterMage,
					Name:         "Fairy",
					Avatar:       record.MageAvatarRedFairy,
					Strength:     10,
					Dexterity:    14,
					Intelligence: 14,
					Coins:        50,
				},
			},
			// Mage - Red Stripe Necromancer
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterMage,
					Name:         "Necromancer",
					Avatar:       record.MageAvatarRedStripeNecromancer,
					Strength:     14,
					Dexterity:    10,
					Intelligence: 14,
					Coins:        50,
				},
			},
			// Mage - Green Elven
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterMage,
					Name:         "Elven",
					Avatar:       record.MageAvatarGreenElven,
					Strength:     12,
					Dexterity:    14,
					Intelligence: 12,
					Coins:        50,
				},
			},
			// Familliar - Brown Cyclops Bat
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterFamilliar,
					Name:         "Brown Cyclops Bat",
					Avatar:       record.FamilliarAvatarBrownCyclopsBat,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Brown Yeti
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterFamilliar,
					Name:         "Brown Yeti",
					Avatar:       record.FamilliarAvatarBrownYeti,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Green Tribble
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterFamilliar,
					Name:         "Green Tribble",
					Avatar:       record.FamilliarAvatarGreenTribble,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Grey Cyclops
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterFamilliar,
					Name:         "Grey Cyclops",
					Avatar:       record.FamilliarAvatarGreyCyclops,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Orange Spotted Tribble
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterFamilliar,
					Name:         "Orange Spotten Tribble",
					Avatar:       record.FamilliarAvatarOrangeSpottedTribble,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Purple Bat
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterFamilliar,
					Name:         "Purple Bat",
					Avatar:       record.FamilliarAvatarPurpleBat,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Purple Minotaur
			{
				Record: record.Entity{
					EntityType:   record.EntityTypeStarterFamilliar,
					Name:         "Purple Minotaur",
					Avatar:       record.FamilliarAvatarPurpleMinotaur,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
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
