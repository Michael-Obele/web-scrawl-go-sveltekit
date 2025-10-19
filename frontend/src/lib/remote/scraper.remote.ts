import { form } from "$app/server";
import * as v from "valibot";
import type { ScrapeResult } from "$lib/types/scraper";

/**
 * Schema for scrape input validation
 */
const ScrapeInputSchema = v.object({
  url: v.pipe(
    v.string(),
    v.nonEmpty("URL is required"),
    v.url("Please enter a valid URL"),
  ),
  depth: v.pipe(
    v.string(),
    v.transform(Number),
    v.number(),
    v.minValue(1, "Depth must be at least 1"),
    v.maxValue(3, "Depth cannot exceed 3"),
  ),
});

/**
 * Remote form function to scrape a website
 * Calls the Go backend API at http://localhost:8080/scrape
 * Works with progressive enhancement - no JavaScript required!
 */
export const scrapeWebsite = form(ScrapeInputSchema, async (data) => {
  // Call the backend API
  const baseUrl = import.meta.env.VITE_API_URL || "http://localhost:8080";
  const apiUrl = `${baseUrl}/scrape?url=${encodeURIComponent(
    data.url,
  )}&depth=${data.depth}`;

  const response = await fetch(apiUrl, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    const errorText = await response.text();
    console.error("❌ Backend request failed:", errorText);
    throw new Error(`Failed to scrape: ${errorText}`);
  }

  console.log(
    "📡 Backend response status:",
    response.status,
    response.statusText,
  );
  console.log(
    "📡 Backend response headers:",
    Object.fromEntries(response.headers.entries()),
  );

  const responseText = await response.text();
  console.log("📄 Raw response text:", responseText);

  let result: ScrapeResult;
  try {
    result = JSON.parse(responseText);
    console.log("✅ Scrape result parsed successfully:", result);
  } catch (error) {
    console.error("❌ Failed to parse JSON:", error);
    throw new Error("Failed to parse response from server.");
  }

  return {
    success: true,
    final: result,
    ...data,
  };
});
