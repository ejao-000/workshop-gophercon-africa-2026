# Sighting Log — design

A calm field guide to wildlife sightings. Read this before changing the look;
keep new pages and components to it unless the attendee asks for something new.

## Palette

Set as custom properties on `:root` in `web/assets/css/app.css`.

| Name | Value | Used for |
|------|-------|----------|
| Paper | `#f6f8f3` | page background |
| Panel | `#ffffff` | cards, tiles, chart cards |
| Tint | `#e9f4ea` | soft green fills, image placeholders |
| Tint strong | `#d7ebdb` | hover borders, the big 404 numeral |
| Ink | `#12251a` | body text |
| Ink soft | `#4d5f53` | secondary text, captions |
| Line | `#dbe6dd` | borders and rules |
| Green | `#15803d` | primary: links, figures, sparkline, markers |
| Green deep | `#0f5132` | headings, buttons, header/footer background |
| Green soft | `#22a150` | photo-less tile initials |
| Amber | `#d97706` | accents, graphics, focus ring, the globe's points |
| Amber ink | `#b45309` | amber used as *text* (the eyebrow label), for contrast |
| Amber soft | `#fbeccb` | soft amber fills |

Contrast: body and secondary text clear 4.5:1 on paper. Amber is only used as
text through `--amber-ink`; `--amber` is for graphics and large marks.

## Lettering

- Headings and figures: **Lora** (serif — "like a field notebook").
- Body and controls: **Work Sans** (clean sans).

Loaded from Google Fonts in `web/views/layout/layout.templ`. `--font-display`
and `--font-body` hold them.

## Shape and space

- Radii: `--radius-sm: 6px`, `--radius: 12px`, `--radius-lg: 22px`; chips and
  buttons are full pills.
- Room, not packed: sections are `clamp(2.5rem, 1.5rem + 4vw, 5rem)` apart.
- Content sits in `.wrap`, a `74rem` column with a fluid gutter.
- Shadows are soft and low (`--shadow-sm`, `--shadow`).

## Components

- **Header**: sticky, translucent paper, logo mark (a compass) + wordmark, a
  pill nav that marks the open page with `aria-current`.
- **Hero**: two columns — the title and a lead on the left, the 3D globe in a
  round tinted panel on the right. Stacks on phones with the globe first.
- **Tiles (the photo wall)**: white cards, a 4:3 photo on top (`object-fit:
  cover`), then the animal's name (linked to its page), its scientific name,
  place, date and the weather. Hover lifts the card and eases the photo. A
  sighting with no photo shows its initial on a striped green field.
- **Filter**: pill chips for animal and country, plus a search box, above the
  wall. Client-side; nothing is removed from the page.
- **Charts**: a green sparkline of sightings per year, and green bars for the
  top animals and countries.
- **Map**: Leaflet on OpenStreetMap, clustered, marker green, radius
  `--radius-lg`.
- **Weather**: a small icon, temperature, description and a wind arrow, on
  every tile and on the detail page.
- **Lightbox**: click any photo everywhere to open it large.
- **Reveal**: sections fade up once as they scroll in.
- **3D globe**: one scene on the home page (`scenes/globe.js`), a green sphere
  with an amber point per sighting. Decoration only; the map and list carry
  the same information.

## Pages

- `/` home: hero, record figures, the photo wall with search, charts, links on.
- `/map` every sighting pinned, clustered.
- `/gallery` every photograph, click to enlarge.
- `/stats` the tallies, charts and leaderboards.
- `/sightings/{id}` one sighting: hero photo, all photos, facts, weather, map.
- 404 page for anything else.

## Motion

Short and calm: 150–200ms hovers, a 0.6s reveal, and nothing that moves on its
own except the globe. Everything respects `prefers-reduced-motion`.

## Avoid

- No emoji as icons; use inline SVG (the tile's pin and calendar, the weather
  sprite).
- Never build an inline style or HTML from GBIF text; `templ.URL` the links and
  keep scripts using `textContent`.
