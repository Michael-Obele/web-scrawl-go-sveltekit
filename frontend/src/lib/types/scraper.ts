// TypeScript types for ScrapeResult (matches backend contract)
export interface Link {
  text: string;
  href: string;
}

export interface ScrapeResult {
  title: string;
  markdown: string;
  links: Link[];
  warnings?: string[];
  fetchedAt: string;
}
