# PolySciChart Requirements

## Application Name
- **Service Name:** PolySciChart
- **Domain:** polyscichart.com
- ** X Account:** PolySciChart 

## Overview
PolySciChart is a service that delivers graphical charts (e.g., line graphs, bar charts, Pareto charts, pie charts,)
showcase **socio-economic data with political context**—including financial markets, government spending,
national GDP, demographic data, industry inputs, national assets, and debt—based on a unique chart ID. The charts
will include labels, legends, and reference links, optimized for rendering in X posts, with revenue generated
primarily through an X account where charts are published.

Content will be optimized for X posts, with a focus on mobile portrait orientation. 

Revenue will be generated primarily through the X account where charts are published and maintained.

The polyscichart.com domain will redirect to the x.com/polyscichart account. With the 
following URL redirect mapping:
- `https://polyscichart.com/chart?id=<chart-id>` will redirect to the X post URL for the chart with the 
  corresponding `<chart-id>`.


---

## Functional Requirements

### 1. Input Processing
- User input and requests is primarily via the X platform.
- Secondary access will be polyscichart.com URL access which redirects to x.com/polyscichart.


### 2. Chart Types
- Support multiple chart types tailored to socio-economic data with political context:
    - **Line graph**: For trends over time (e.g., national debt growth by administration).
    - **Bar chart**: For comparisons (e.g., industry outputs under different tax regimes).
    - **Pareto chart**: For prioritizing factors (e.g., top demographic contributors to GDP).
    - **Pie chart**: For proportional breakdowns (e.g., government spending by sector).
    - **Scatter plot**: For correlations (e.g., stock market performance vs. political stability).
    - **Bubble chart**: For multi-dimensional analysis (e.g., state GDP vs. population vs. party control).
- Chart type is predefined and stored with the `<chart-id>` data (e.g., a metadata field like `"type": "bubble"`).

### 3. Chart Features
- Include minimal predefined labels (e.g., numeric ticks) or no text, relying on X post text/alt text for context, 
based on socio-economic-political data tied to the `<chart-id>`.
- Use color, shape, or size to identify data series or categories (e.g., political parties, fiscal years), minimizing 
or omitting legends.
- Exclude embedded sponsor messages or reference links from the image, placing them in the X post text instead.
- Ensure charts are visually clean, high-fidelity, and readable in small sizes (optimized for X posts).

### 4. Output
- Generate X posts
  - Generate charts as PNG image files for easy embedding in X posts.
    - images should be optimized for mobile portrait view on modern mobile phone devices
  - Generate X posts with the chart image (PNG) plus text description, hashtags, and in future sponsor.
- Preserve source polyscichart-post data on a private S3 object store on the OCI service

---

---

## PolySciChart chart manager Workflow Example
1. User creates PSCPost source data and uses a tool to generate the post data and store to a private object store.
2. User previews chart generations and X post data.
3. User posts the X post.

## End Consumer Workflow Example
1. User 
   2. follows the X account PolySciChart and receives a post in their feed.
   3. discovers a chart post from a friend or follower.
   4. discovers a chart post from a URL link in some other source.
5. User clicks on the chart post and views the chart.

---

## Monetization Strategy
- **Primary Revenue:** Generated via the PolySciChart X account through:
    - Ad share from traffic view PolySciChart posts.
- Leverage X’s ad revenue tools (e.g., X Premium or promoted posts) to amplify reach and earnings.

---

## Branding Notes
- "PolySciChart" reflects its focus on **socio-economic data with political context**, combining "PolySci" 
  (Political Science) with "Chart" (data visualization). It delivers charts that blend economic metrics—like 
  financial markets, government spending, national GDP, industry inputs, national assets, and debt—with social 
  factors (e.g., demographics) and political contexts (e.g., party control, policy impacts, electoral cycles).
- Example charts could include:
    - Government spending by party (pie chart).
    - Financial market trends during election years (scatter plot).
    - National debt vs. demographic voting patterns vs. political leadership (bubble chart).
- The sponsor message and branding should remain subtle and professional (e.g., small font, bottom-right placement).