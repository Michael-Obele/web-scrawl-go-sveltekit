// @ts-check

/** @type {import('./$types').PageServerLoad} */
export async function load({ fetch }) {
	try {
		const response = await fetch('http://localhost:8080/health');
		
		if (!response.ok) {
			return {
				health: null,
				error: 'Backend service is not available'
			};
		}

		const health = await response.json();
		return { health, error: null };
	} catch (error) {
		return {
			health: null,
			error: 'Backend service is not available'
		};
	}
}
