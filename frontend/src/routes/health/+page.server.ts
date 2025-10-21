import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  const startTime = Date.now();

  // Determine runtime mode (development|production) and backend base URL
  const mode = (import.meta.env.MODE as string) || "development";
  const baseUrl = import.meta.env.VITE_API_URL || "http://localhost:8080";

  // Explain what the UI will do to try to wake the backend
  const wakeInstructions =
    mode === "production"
      ? `This page will attempt to open the backend URL (${baseUrl}) in a new tab and poll the /health endpoint until it responds. Hosted environments may take several seconds to wake.`
      : `On your local machine this will open the backend URL (${baseUrl}) in a new tab. If you're developing locally, start the Go backend: cd backend && go run main.go`;

  try {
    const response = await fetch(`${baseUrl}/health`);
    const duration = Date.now() - startTime;

    if (!response.ok) {
      return {
        health: null,
        error: "Backend service is not available",
        duration,
        env: mode,
        wakeInstructions,
      };
    }

    const health = await response.json();
    console.log({ health, error: null, duration, env: mode, wakeInstructions });
    return { health, error: null, duration, env: mode, wakeInstructions };
  } catch (error) {
    const duration = Date.now() - startTime;
    console.error({ error: null, duration, env: mode, wakeInstructions });
    return {
      health: null,
      error: `Backend service is not available, ${error}`,
      duration,
      env: mode,
      wakeInstructions,
    };
  }
};
