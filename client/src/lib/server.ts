import { HealthReponse } from './types/responses/responses';

export const getServerHealth = async (): Promise<HealthReponse> => {
  const res = await fetch('http://localhost:8080/health', {
    next: { revalidate: 60 },
  });
  if (!res.ok) throw new Error('Health check failed');
  const finalRes: HealthReponse = await res.json();
  return finalRes;
};
