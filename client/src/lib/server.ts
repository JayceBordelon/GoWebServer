import { HealthReponse } from './types/responses/responses';

export const getServerHealth = async (): Promise<HealthReponse> => {
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_BASE_SERVER_URL}/health`, {
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
