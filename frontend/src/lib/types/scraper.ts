// TypeScript types for ScrapeResult (matches backend contract)
export interface Link {
  text: string;
  href: string;
}

export interface ScrapeResult {
  title: string;
  markdown: string;
  links: Link[];
  rawHtml?: string;
  warnings?: string[];
  fetchedAt: string;
}

// Types for the remote function return values
export interface ScrapeSuccessResult {
  url: string;
  depth: number;
  success: true;
  final: ScrapeResult;
}

export interface ScrapeErrorResult {
  url: string;
  depth: number;
  success: false;
  error: string;
}

export type ScrapeFormResult = ScrapeSuccessResult | ScrapeErrorResult;
