import { DateTime } from "luxon";

export default function(eleventyConfig) {
	// Date filter
	eleventyConfig.addFilter("readableDate", (dateObj) => {
		return DateTime.fromJSDate(dateObj, { zone: "utc" }).toFormat(
			"dd LLLL yyyy"
		);
	});

	eleventyConfig.addNunjucksAsyncFilter("readableDateAsync", (dateObj, callback) => {
		try {
			const date = DateTime.fromJSDate(dateObj, { zone: "utc" }).toFormat(
				"dd LLLL yyyy"
			);
			callback(null, date);
		} catch (err) {
			callback(err);
		}
	});

	// Date filter for year
	eleventyConfig.addFilter("year", () => `${new Date().getFullYear()}`);

	// Min filter
	eleventyConfig.addFilter("min", (...numbers) => {
		return Math.min.apply(null, numbers);
	});

	// Concat filter
	eleventyConfig.addFilter("concat", function(value) {
		return value.join("");
	});
}
