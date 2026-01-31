package lint

import "unagent-lint/rules"

func ActiveRules(cfg Config) []rules.Rule {
	r := []rules.Rule{
		rules.AccessibilityFraming(),
		rules.NonDeficitLanguage(),
		rules.NotAnAgent(),
		rules.NoImpersonation(),
		rules.NoDeception(),
		rules.ConsentBoundary(),
		rules.CognitiveLoadReduction(),
		rules.DefaultsFirst(),
		rules.OfflineSafety(),
	}

	if !cfg.NoVoice {
		r = append(r,
			rules.CalmAuthority(),
			rules.EmotionalTemperature(),
			rules.PrecisionOverVolume(),
		)
	}
	return r
}
