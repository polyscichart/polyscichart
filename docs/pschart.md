# X Post CSV Template for @polyscichart

This document defines the CSV template used to create X posts for the `@polyscichart` account, including styled text and chart images. The template is processed by Go utilities (`checkcsv.go` and `postx.go`) to generate and post content to X, then record the results.

## Structure
The CSV is divided into three sections with unique delimiters:
- **Header Section**: Input fields for the X post and chart, before the `+++` delimiter.
- **Generated Section**: Fields written after posting, between `+++` and `---`.
- **Chart Data Section**: Data for chart generation, after `---`.

Each row uses column A as the key and columns B, C, D (etc.) as values, with commas separating fields.

---

## Example: Before Posting (`post1.csv`)

```aiignore
Key, Value1, Value2, Value3
x_title,"US GDP Growth 2024",,
x_text,"Q1 up 2.1% #GDP, Q2 at 1.9%, Q3 at 2.5% @xAI
",,
x_sponsor,"",,
x_source,"FRED",,
chart_type,bar,,
image_path,,,
title,"GDP Growth 2024",,
chart_alt,"Bar chart of US GDP growth rates for 2024: Q1 2.1%, Q2 1.9%, Q3 2.5%",,
style,"bgcolor=#f0f0f0","color=#3366cc",,
+++
Year,Q1,Q2,Q3
2024,2.1,1.9,2.5


```

## Example: After Posting (`post1.csv`)

```aiignore
Key, Value1, Value2, Value3
x_title,"US GDP Growth 2024",,
x_text,"Q1 up 2.1% #GDP, Q2 at 1.9%, Q3 at 2.5% @xAI
",,
x_sponsor,"",,
x_source,"FRED",,
chart_type,bar,,
image_path,,,
title,"GDP Growth 2024",,
chart_alt,"Bar chart of US GDP growth rates for 2024: Q1 2.1%, Q2 1.9%, Q3 2.5%",,
style,"bgcolor=#f0f0f0","color=#3366cc",,
+++
x_post_id,"1234567890123456789",,
x_post_url,"https://x.com/polyscichart/status/1234567890123456789",,
x_media_id,"9876543210987654321",,
x_created_at,"2025-03-15T12:34:56Z",,
Year,Q1,Q2,Q3
2024,2.1,1.9,2.5

```

## Resulting X Post
- **Text**: **US GDP Growth 2024**. Q1 up 2.1% #GDP, Q2 at 1.9%, Q3 at 2.5% @xAI. *FRED*
  - Length: 92 characters (under 280).
- **Image**: Bar chart with title "GDP Growth 2024", styled with `bgcolor=#f0f0f0` and `color=#3366cc`.
- **Alt Text**: "Bar chart of US GDP growth rates for 2024: Q1 2.1%, Q2 1.9%, Q3 2.5%"

---

## Field Explanations

### Header Section (Before `+++`)
These fields define the X post content and chart generation. All are optional unless specified.

| Key          | Description                                                                                   | Example Value                          | Notes                                                                                   |
|--------------|-----------------------------------------------------------------------------------------------|----------------------------------------|-----------------------------------------------------------------------------------------|
| `x_title`    | Tweet title, displayed in **bold**.                                                          | "US GDP Growth 2024"                  | Required for the tweet; prefixed to the final text.                                     |
| `x_text`     | Main tweet body, plain text. Can include hashtags (e.g., `#GDP`) and mentions (e.g., `@xAI`). | "Q1 up 2.1% #GDP, Q2 at 1.9%, ..."   | Required; forms the core of the tweet. No carriage returns allowed (stripped if present). |
| `x_sponsor`  | Sponsor name, displayed in ***italic-bold*** if provided.                                    | "xAI" or ""                            | Optional; appended after `x_text` with a space if non-empty.                            |
| `x_source`   | Source(s) of data, displayed in *italic*. Multiple values joined with `. `.                  | "FRED" or "Census","World Bank"       | Optional; appended after `x_sponsor` with `. `. Multiple sources supported across columns. |
| `chart_type` | Type of chart to generate (e.g., `bar`, `pie`). Ignored if `image_path` is set.              | "bar" or ""                            | Required for chart generation if `image_path` is empty.                                 |
| `image_path` | Path to an existing PNG file. If empty, a chart is generated using `chart_type`.            | "charts/pop_state.png" or ""           | Optional; overrides chart generation if set.                                            |
| `title`      | Title for the generated chart. Ignored if `image_path` is set.                               | "GDP Growth 2024"                      | Optional; used by `go-chart` for the chart’s title.                                     |
| `chart_alt`  | Alt text for the chart image (max 1000 chars). Applied to generated or uploaded images.      | "Bar chart of US GDP growth rates..." | Optional; enhances accessibility via X’s alt text feature.                              |
| `style`      | Styling attributes for the chart (e.g., `bgcolor`, `color`). Multiple values across columns. | "bgcolor=#f0f0f0","color=#3366cc"     | Optional; `key=value` pairs (e.g., `bgcolor`, `color`) for `go-chart` styling.          |

- **Delimiter**: `+++` marks the end of input fields and start of generated fields.

### Generated Section (Between `+++` and `---`)
These fields are written by `postx.go` after a successful X post.

| Key           | Description                                                                                 | Example Value                                    | Notes                                                                                     |
|---------------|---------------------------------------------------------------------------------------------|--------------------------------------------------|-------------------------------------------------------------------------------------------|
| `x_post_id`   | Tweet ID returned by the X API v2 `POST /2/tweets`.                                        | "1234567890123456789"                           | Unique identifier for the posted tweet.                                           |
| `x_post_url`  | URL to the tweet, constructed as `https://x.com/polyscichart/status/<x_post_id>`.          | "https://x.com/polyscichart/status/1234567890" | Direct link to the post; username is hardcoded as `polyscichart`.                        |
| `x_media_id`  | Media ID of the uploaded chart image, from v1.1 `POST media/upload`.                       | "9876543210987654321"                           | Identifier for the chart image in X’s system.                                            |
| `x_created_at`| Creation timestamp of the tweet (ISO 8601), from X API or generated locally if unavailable.| "2025-03-15T12:34:56Z"                          | Reflects when the tweet was posted; prefers X API value if provided.                     |

- **Delimiter**: `---` marks the end of generated fields and start of chart data.

### Chart Data Section (After `---`)
Data used to generate the chart if `image_path` is empty.

| Row Type | Description                                   | Example Value            | Notes                                                  |
|----------|-----------------------------------------------|--------------------------|--------------------------------------------------------|
| Labels   | First row: Labels for chart categories/axes. | "Year,Q1,Q2,Q3"         | Defines the X-axis or categories (e.g., quarters).     |
| Values   | Subsequent rows: Numeric data for the chart. | "2024,2.1,1.9,2.5"      | Data points; one row for simplicity in this example.  |

- **Requirement**: Only used when `image_path` is empty and `chart_type` is set.

---

## Tweet Construction
The final tweet text is built as:
- **Format**: **`x_title`**. `x_text` ***x_sponsor***. *x_source*
- **Rules**:
  - `. ` separates `x_title` and `x_text`.
  - Space before `x_sponsor` if non-empty.
  - `. ` before `x_source` if non-empty; multiple sources joined with `. `.
  - Must fit within 280 characters (Basic Premium limit).
- **Example**: **US GDP Growth 2024**. Q1 up 2.1% #GDP, Q2 at 1.9%, Q3 at 2.5% @xAI. *FRED*

## Notes
- **Premium Styling**: Requires a Basic Premium X account for native bold and italic formatting.
- **Excel Compatibility**: `+++` and `---` are treated as text, not affecting basic editing or charting.
- **Chart Generation**: Uses `go-chart` with `chart_type`, `title`, and `style` if `image_path` is empty.