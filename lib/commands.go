package lib

import (
	"strconv"
)

// Calls the API and prints every website in a table
func WebsitesCommand() {
	res := WebsiteResponse{}
	apiGet(GetConfiguration().ApplicationUrl+"/api/v1/websites", &res)

	if res.Success && len(res.Websites) >= 1 {
		// put results into two-dimensional array
		data := make([][]string, len(res.Websites))
		for i := range data {
			data[i] = make([]string, 3)
			data[i][0] = res.Websites[i].Name
			data[i][1] = res.Websites[i].Protocol + "://" + res.Websites[i].Url
			data[i][2] = res.Websites[i].Status
		}

		// print as table
		PrintAsTable([]string{"Name", "URL", "Status"}, data)
	} else {
		println("Nothing found.")
	}
}

// Calls the API and prints the given website's status-information
func StatusCommand(website string) {
	res := StatusResponse{}
	apiGet(GetConfiguration().ApplicationUrl+"/api/v1/websites/"+website+"/status", &res)

	if res.Success {
		println(res.WebsiteData.Name + " (" + res.WebsiteData.Url + ")")
		println("Status: " + res.LastCheckResult.Status + " (" + res.LastCheckResult.ResponseTime + ") @ " + res.LastCheckResult.Time)
		if res.LastFailedCheckResult.Time != "0000-00-00 00:00:00" {
			println("Last Fail: " + res.LastFailedCheckResult.Status + " (" + res.LastFailedCheckResult.ResponseTime + ") @ " + res.LastFailedCheckResult.Time)
		}
		println("Availability: " + res.Availability.Average + " (" + strconv.Itoa(res.Availability.Total) + " Checks)")
	} else {
		println("Nothing found.")
	}
}

// Calls the API and prints the given website's results
func ResultsCommand(website string, limit int, offset int) {
	if limit > 9999 || limit < 0 || offset > 9999 || offset < 0 {
		PrintWrongArguments()
	}

	res := ResultsResponse{}
	apiGet(GetConfiguration().ApplicationUrl+"/api/v1/websites/"+website+"/results?limit="+strconv.Itoa(limit)+"&offset="+strconv.Itoa(offset), &res)

	if res.Success {
		// put results into two-dimensional array
		data := make([][]string, len(res.Websites))
		for i := range data {
			data[i] = make([]string, 3)
			data[i][0] = res.Websites[i].Time
			data[i][1] = res.Websites[i].Status
			data[i][2] = res.Websites[i].ResponseTime
		}

		// print as table
		PrintAsTable([]string{"Time", "Status", "Response Time"}, data)
	} else {
		println("Nothing found.")
	}
}
