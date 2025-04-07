import { getServerHealth } from '@lib/server';

export default async function Home() {
  let data = null;
  try {
    data = await getServerHealth();
  } catch (err) {
    console.warn('Health check failed during build:', err);
  }

  return (
    <div>
      <main>{JSON.stringify(data || { status: 'unreachable' })}</main>
    </div>
  );
}
