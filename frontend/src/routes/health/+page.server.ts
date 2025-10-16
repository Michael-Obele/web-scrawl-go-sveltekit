import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  try {
    const response = await fetch("http://localhost:8080/health");

    if (!response.ok) {
      return {
        health: null,
        error: "Backend service is not available",
      };
    }

    const health = await response.json();
    return { health, error: null };
  } catch (error) {
    return {
      health: null,
      error: "Backend service is not available",
    };
  }
};
