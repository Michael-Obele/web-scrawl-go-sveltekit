import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  const startTime = Date.now();
  try {
    const baseUrl = import.meta.env.VITE_API_URL || "http://localhost:8080";
    const response = await fetch(`${baseUrl}/health`);
    const duration = Date.now() - startTime;

    if (!response.ok) {
      return {
        health: null,
        error: "Backend service is not available",
        duration,
      };
    }

    const health = await response.json();
    return { health, error: null, duration };
  } catch (error) {
    const duration = Date.now() - startTime;
    return {
      health: null,
      error: "Backend service is not available",
      duration,
    };
  }
};
