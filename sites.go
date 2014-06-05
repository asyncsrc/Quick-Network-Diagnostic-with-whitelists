package main

func populateSites() map[string][]string {
	sites := make(map[string][]string)

	cloudFrontSites := [...]string{
		"https://s3.amazonaws.com/",
		"https://url_here",
	}
	sites["Cloud Front Sites"] = cloudFrontSites[:]

	requiredSites := [...]string{
		"https://www.pfrsolutions.com",
		"https://url_here",
		"https://url_here",
	}
	sites["Required Sites"] = requiredSites[:]

	learnContentSites := [...]string{
		"https://s3.amazonaws.com/xyz",
		"https://url_here",
		"https://url_here",
	}
	sites["Learn Content Sites"] = learnContentSites[:]

	showMeSites := [...]string{
		"https://s3.amazonaws.com/s3.maketutorial.com/users",
	}
	sites["ShowMe Sites"] = showMeSites[:]

	ie8RequiredSites := [...]string{
		"https://url_here",
	}
	sites["IE8 Required Sites"] = ie8RequiredSites[:]

	optionalSites := [...]string{
		"https://url_here",
		"https://url_here",
		"https://url_here",
		"https://url_here",
		"https://url_here",
		"https://url_here",
		"https://url_here",
	}
	sites["Optional Sites"] = optionalSites[:]

	mobileSites := [...]string{
		"https://url_here",
		"https://url_here",
	}
	sites["Mobile Sites"] = mobileSites[:]

	backgroundCheckSites := [...]string{
		"http://www.irs.gov/",
	}
	sites["Background Check Sites"] = backgroundCheckSites[:]

	applicantSites := [...]string{
		"https://url_here",
		"https://url_here",
	}
	sites["Applicant Sites"] = applicantSites[:]

	supportSites := [...]string{
		"https://url_here",
		"https://url_here",
		"https://url_here",
	}
	sites["Support Sites"] = supportSites[:]

	return sites
}
