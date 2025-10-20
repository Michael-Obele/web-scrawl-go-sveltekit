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
    v.url("Please enter a valid URL")
  ),
  depth: v.optional(
    v.pipe(
      v.number(),
      v.minValue(1, "Depth must be at least 1"),
      v.maxValue(3, "Depth cannot exceed 3")
    ),
    1 // Default value of "1" (will be transformed to number 1)
  ),
});

/**
 * Remote form function to scrape a website
 * Calls the Go backend API at http://localhost:8080/scrape
 * Works with progressive enhancement - no JavaScript required!
 */
export const scrapeWebsite = form(ScrapeInputSchema, async (data) => {
  // Check backend health first
  const baseUrl = import.meta.env.VITE_API_URL || "http://localhost:8080";

  try {
    const healthResponse = await fetch(`${baseUrl}/health`);
    if (!healthResponse.ok) {
      throw new Error("Backend service is not responding");
    }
  } catch (error) {
    console.error("❌ Backend health check failed:", error);
    throw new Error(
      "Backend is sleeping. Please visit the health page to wake it up."
    );
  }

  // Call the backend API
  const apiUrl = `${baseUrl}/scrape?url=${encodeURIComponent(data.url)}&depth=${
    data.depth
  }`;

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
    response.statusText
  );
  console.log(
    "📡 Backend response headers:",
    Object.fromEntries(response.headers.entries())
  );

  const responseText = await response.text();
  // console.log("📄 Raw response text:", responseText);

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
