package features_test

import (
	"github.com/wfernandes/weather/wunderground/features"

	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Conditions", func() {

	It("current_observation JSON body can be unmarshalled", func() {

		jsonStr := `{"image":{"url":"http://icons.wxug.com/graphics/wu2/logo_130x80.png","title":"Weather Underground","link":"http://www.wunderground.com"},"display_location":{"full":"Boulder, CO","city":"Boulder","state":"CO","state_name":"Colorado","country":"US","country_iso3166":"US","zip":"80301","magic":"1","wmo":"99999","latitude":"40.04679489","longitude":"-105.21250153","elevation":"1605.00000000"},"observation_location":{"full":"63rd St and Jay Rd, Boulder, Colorado","city":"63rd St and Jay Rd, Boulder","state":"Colorado","country":"US","country_iso3166":"US","latitude":"40.055550","longitude":"-105.208595","elevation":"5202 ft"},"estimated":{},"station_id":"KCOBOULD39","observation_time":"Last Updated on June 10, 9:47 PM MDT","observation_time_rfc822":"Wed, 10 Jun 2015 21:47:27 -0600","observation_epoch":"1433994447","local_time_rfc822":"Wed, 10 Jun 2015 21:47:59 -0600","local_epoch":"1433994479","local_tz_short":"MDT","local_tz_long":"America/Denver","local_tz_offset":"-0600","weather":"Overcast","temperature_string":"64.2 F (17.9 C)","temp_f":64.2,"temp_c":17.9,"relative_humidity":"80%","wind_string":"Calm","wind_dir":"North","wind_degrees":354,"wind_mph":0,"wind_gust_mph":"0","wind_kph":0,"wind_gust_kph":"0","pressure_mb":"1011","pressure_in":"29.87","pressure_trend":"-","dewpoint_string":"58 F (14 C)","dewpoint_f":58,"dewpoint_c":14,"heat_index_string":"NA","heat_index_f":"NA","heat_index_c":"NA","windchill_string":"NA","windchill_f":"NA","windchill_c":"NA","feelslike_string":"64.2 F (17.9 C)","feelslike_f":"64.2","feelslike_c":"17.9","visibility_mi":"10.0","visibility_km":"16.1","solarradiation":"--","UV":"0","precip_1hr_string":"0.00 in ( 0 mm)","precip_1hr_in":"0.00","precip_1hr_metric":" 0","precip_today_string":"0.00 in (0 mm)","precip_today_in":"0.00","precip_today_metric":"0","icon":"cloudy","icon_url":"http://icons.wxug.com/i/c/k/nt_cloudy.gif","forecast_url":"http://www.wunderground.com/US/CO/Boulder.html","history_url":"http://www.wunderground.com/weatherstation/WXDailyHistory.asp?ID=KCOBOULD39","ob_url":"http://www.wunderground.com/cgi-bin/findweather/getForecast?query=40.055550,-105.208595","nowcast":""}`

		var cond features.Conditions
		err := json.Unmarshal([]byte(jsonStr), &cond)
		Expect(err).ToNot(HaveOccurred())

		Expect(cond.Image.Url).To(Equal("http://icons.wxug.com/graphics/wu2/logo_130x80.png"))
		Expect(cond.Image.Title).To(Equal("Weather Underground"))
		Expect(cond.Image.Link).To(Equal("http://www.wunderground.com"))

		Expect(cond.DisplayLocation.Zip).To(Equal("80301"))
		Expect(cond.DisplayLocation.City).To(Equal("Boulder"))
		Expect(cond.DisplayLocation.Country).To(Equal("US"))
		Expect(cond.DisplayLocation.Elevation).To(Equal("1605.00000000"))
		Expect(cond.DisplayLocation.Latitude).To(Equal("40.04679489"))
		Expect(cond.DisplayLocation.Longitude).To(Equal("-105.21250153"))
		Expect(cond.DisplayLocation.State).To(Equal("CO"))
		Expect(cond.DisplayLocation.StateName).To(Equal("Colorado"))

		Expect(cond.ObservationLocation.City).To(Equal("63rd St and Jay Rd, Boulder"))
		Expect(cond.ObservationLocation.Country).To(Equal("US"))
		Expect(cond.ObservationLocation.Elevation).To(Equal("5202 ft"))
		Expect(cond.ObservationLocation.Latitude).To(Equal("40.055550"))
		Expect(cond.ObservationLocation.Longitude).To(Equal("-105.208595"))
		Expect(cond.ObservationLocation.State).To(Equal("Colorado"))

		Expect(cond.StationId).To(Equal("KCOBOULD39"))
		Expect(cond.ObservationTimeRFC822).To(Equal("Wed, 10 Jun 2015 21:47:27 -0600"))
		Expect(cond.ObservationEpoch).To(Equal("1433994447"))
		Expect(cond.LocalTimeRFC822).To(Equal("Wed, 10 Jun 2015 21:47:59 -0600"))
		Expect(cond.LocalEpoch).To(Equal("1433994479"))
		Expect(cond.LocalTZShort).To(Equal("MDT"))
		Expect(cond.LocalTZOffset).To(Equal("-0600"))
		Expect(cond.Weather).To(Equal("Overcast"))
		Expect(cond.TempF).To(Equal(64.2))
		Expect(cond.TempC).To(Equal(17.9))
		Expect(cond.RelativeHumidity).To(Equal("80%"))
		Expect(cond.Wind).To(Equal("Calm"))
		Expect(cond.WindDir).To(Equal("North"))
		Expect(cond.WindDegrees).To(BeEquivalentTo(354))
		Expect(cond.WindMPH).To(BeEquivalentTo(0))
		Expect(cond.WindGustMPH).To(BeEquivalentTo("0"))
		Expect(cond.WindKPH).To(BeEquivalentTo(0))
		Expect(cond.WindGustKPH).To(BeEquivalentTo("0"))
		Expect(cond.PressureMB).To(Equal("1011"))
		Expect(cond.PressureIN).To(Equal("29.87"))
		Expect(cond.PressureTrend).To(Equal("-"))
		Expect(cond.DewpointF).To(BeEquivalentTo(58))
		Expect(cond.DewpointC).To(BeEquivalentTo(14))
		Expect(cond.FeelslikeF).To(Equal("64.2"))
		Expect(cond.FeelslikeC).To(Equal("17.9"))
		Expect(cond.VisibilityMI).To(Equal("10.0"))
		Expect(cond.VisibilityKM).To(Equal("16.1"))
		Expect(cond.Precip1hr).To(Equal("0.00 in ( 0 mm)"))
		Expect(cond.Precip1hrIN).To(Equal("0.00"))
		Expect(cond.Precip1hrMETRIC).To(Equal(" 0"))
		Expect(cond.PrecipToday).To(Equal("0.00 in (0 mm)"))
		Expect(cond.PrecipTodayIN).To(Equal("0.00"))
		Expect(cond.PrecipTodayMETRIC).To(Equal("0"))
		Expect(cond.Icon).To(Equal("cloudy"))
		Expect(cond.IconURL).To(Equal("http://icons.wxug.com/i/c/k/nt_cloudy.gif"))
		Expect(cond.ForecastURL).To(Equal("http://www.wunderground.com/US/CO/Boulder.html"))
		Expect(cond.HistoryURL).To(Equal("http://www.wunderground.com/weatherstation/WXDailyHistory.asp?ID=KCOBOULD39"))
		Expect(cond.ObURL).To(Equal("http://www.wunderground.com/cgi-bin/findweather/getForecast?query=40.055550,-105.208595"))

		Expect(cond.Weather).To(Equal("Overcast"))
	})

	It("can be unmarshalled from the full JSON response", func() {

		jsonResp := `{"response":{"version":"0.1","termsofService":"http://www.wunderground.com/weather/api/d/terms.html","features":{"conditions":1}},"current_observation":{"image":{"url":"http://icons-ak.wxug.com/graphics/wu2/logo_130x80.png","title":"Weather Underground","link":"http://www.wunderground.com"},"display_location":{"full":"San Francisco, CA","city":"San Francisco","state":"CA","state_name":"California","country":"US","country_iso3166":"US","zip":"94101","latitude":"37.77500916","longitude":"-122.41825867","elevation":"47.00000000"},"observation_location":{"full":"SOMA - Near Van Ness, San Francisco, California","city":"SOMA - Near Van Ness, San Francisco","state":"California","country":"US","country_iso3166":"US","latitude":"37.773285","longitude":"-122.417725","elevation":"49 ft"},"estimated":{},"station_id":"KCASANFR58","observation_time":"Last Updated on June 27, 5:27 PM PDT","observation_time_rfc822":"Wed, 27 Jun 2012 17:27:13 -0700","observation_epoch":"1340843233","local_time_rfc822":"Wed, 27 Jun 2012 17:27:14 -0700","local_epoch":"1340843234","local_tz_short":"PDT","local_tz_long":"America/Los_Angeles","local_tz_offset":"-0700","weather":"Partly Cloudy","temperature_string":"66.3 F (19.1 C)","temp_f":66.3,"temp_c":19.1,"relative_humidity":"65%","wind_string":"From the NNW at 22.0 MPH Gusting to 28.0 MPH","wind_dir":"NNW","wind_degrees":346,"wind_mph":22.0,"wind_gust_mph":"28.0","wind_kph":35.4,"wind_gust_kph":"45.1","pressure_mb":"1013","pressure_in":"29.93","pressure_trend":"+","dewpoint_string":"54 F (12 C)","dewpoint_f":54,"dewpoint_c":12,"heat_index_string":"NA","heat_index_f":"NA","heat_index_c":"NA","windchill_string":"NA","windchill_f":"NA","windchill_c":"NA","feelslike_string":"66.3 F (19.1 C)","feelslike_f":"66.3","feelslike_c":"19.1","visibility_mi":"10.0","visibility_km":"16.1","solarradiation":"","UV":"5","precip_1hr_string":"0.00 in ( 0 mm)","precip_1hr_in":"0.00","precip_1hr_metric":" 0","precip_today_string":"0.00 in (0 mm)","precip_today_in":"0.00","precip_today_metric":"0","icon":"partlycloudy","icon_url":"http://icons-ak.wxug.com/i/c/k/partlycloudy.gif","forecast_url":"http://www.wunderground.com/US/CA/San_Francisco.html","history_url":"http://www.wunderground.com/history/airport/KCASANFR58/2012/6/27/DailyHistory.html","ob_url":"http://www.wunderground.com/cgi-bin/findweather/getForecast?query=37.773285,-122.417725"  }}`

		var weatherResp features.ConditionsResponse
		err := json.Unmarshal([]byte(jsonResp), &weatherResp)
		Expect(err).ToNot(HaveOccurred())

		Expect(weatherResp.Response).ToNot(BeNil())
		Expect(weatherResp.Response.Error).To(BeNil())
		Expect(weatherResp.Condition).ToNot(BeNil())
		Expect(weatherResp.Condition.Weather).To(Equal("Partly Cloudy"))
	})
})
