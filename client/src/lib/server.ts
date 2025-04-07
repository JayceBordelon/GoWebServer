import { HealthReponse } from './types/responses/responses';

export const getServerHealth = async (): Promise<HealthReponse> => {
  if (process.env.NEXT_PHASE === 'phase-production-build') {
    return { message: 'skipping health check', status: 'skipped (build time)' };
  }

  try {
    const res = await fetch('http://localhost:8080/health', {
      next: { revalidate: 60 },
    });

    if (!res.ok) {
      throw new Error(`Health check failed with status: ${res.status}`);
    }

    const finalRes: HealthReponse = await res.json();
    return finalRes;
  } catch (error) {
    console.warn('Health check failed:', error);
    return { message: 'Failure', status: 'unreachable' };
  }
};
