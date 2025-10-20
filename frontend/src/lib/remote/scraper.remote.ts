import { form } from "$app/server";
import * as v from "valibot";
import type { ScrapeResult, ScrapeFormResult } from "$lib/types/scraper";

/**
 * Schema for scrape input validation
 */
const ScrapeInputSchema = v.object({
  url: v.pipe(
    v.string(),
    v.nonEmpty("URL is required"),
    v.url("Please enter a valid URL")
  ),
  depth: v.fallback(
    v.pipe(
      v.string(),
      v.check((value) => value === "" || !isNaN(Number(value)), "Depth must be a valid number"),
      v.transform((value) => value === "" ? 1 : Number(value)),
      v.number(),
      v.minValue(1, "Depth must be at least 1"),
      v.maxValue(3, "Depth cannot exceed 3")
    ),
    1
  ),
});

/**
 * Remote form function to scrape a website
 * Calls the Go backend API at http://localhost:8080/scrape
 * Works with progressive enhancement - no JavaScript required!
 */
export const scrapeWebsite = form(
  ScrapeInputSchema,
  async (data): Promise<ScrapeFormResult> => {
    // Check backend health first
    const baseUrl = import.meta.env.VITE_API_URL || "http://localhost:8080";

    try {
      const healthResponse = await fetch(`${baseUrl}/health`);
      if (!healthResponse.ok) {
        return {
          success: false,
          error:
            "Backend service is not responding. Please visit the health page to wake it up.",
          url: data.url,
          depth: data.depth,
        } as const;
      }
    } catch (error) {
      console.error("❌ Backend health check failed:", error);
      return {
        success: false,
        error:
          "Backend is sleeping. Please visit the health page to wake it up.",
        url: data.url,
        depth: data.depth,
      } as const;
    }

    // Call the backend API
    const apiUrl = `${baseUrl}/scrape?url=${encodeURIComponent(
      data.url
    )}&depth=${data.depth}`;

    let response: Response;
    try {
      response = await fetch(apiUrl, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });
    } catch (error) {
      console.error("❌ Network error:", error);
      return {
        success: false,
        error: "Network error: Unable to connect to the scraping service.",
        url: data.url,
        depth: data.depth,
      } as const;
    }

    if (!response.ok) {
      let errorMessage = "Unknown error occurred";
      try {
        const errorData = await response.json();
        errorMessage =
          errorData.message ||
          errorData.error ||
          `HTTP ${response.status}: ${response.statusText}`;
      } catch {
        // If we can't parse the error response, use the status text
        errorMessage = `HTTP ${response.status}: ${response.statusText}`;
      }

      console.error("❌ Backend request failed:", errorMessage);
      return {
        success: false,
        error: `Scraping failed: ${errorMessage}`,
        url: data.url,
        depth: data.depth,
      } as const;
    }

    console.log(
      "📡 Backend response status:",
      response.status,
      response.statusText
    );

    const responseText = await response.text();

    let result: ScrapeResult;
    try {
      result = JSON.parse(responseText);
      console.log("✅ Scrape result parsed successfully:", result);
    } catch (error) {
      console.error("❌ Failed to parse JSON:", error);
      return {
        success: false,
        error:
          "Failed to parse response from server. The scraping service may have returned invalid data.",
        url: data.url,
        depth: data.depth,
      } as const;
    }

    return {
      success: true,
      final: result,
      url: data.url,
      depth: data.depth,
    } as const;
  }
);
