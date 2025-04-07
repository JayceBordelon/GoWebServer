import { getServerHealth } from '@lib/server';
import { Container, Title, Text, Button, Stack, Group, Image, Paper, Badge } from '@mantine/core';
import { IconBrandGolang, IconCloudQuestion, IconRocket } from '@tabler/icons-react';

export default async function Home() {
  let data = null;
  try {
    data = await getServerHealth();
    console.log(data);
  } catch (err) {
    console.warn('Health check failed during build:', err);
  }

  return (
    <Container size="lg" py="xl">
      <Stack gap="xl" align="center">
        <Stack gap={8} align="center" ta="center">
          <Title order={1} size="h2">
            Effortless Image Hosting
          </Title>
          <Text size="lg" c="dimmed" maw={600}>
            Store, serve, and transform images in seconds — no config, no stress.
          </Text>
        </Stack>

        <Group mt="sm" justify="center">
          <Button size="md" variant="outline" leftSection={<IconCloudQuestion size={20} />}>
            How it works
          </Button>
          <Button size="md" leftSection={<IconRocket size={20} />}>
            Get started
          </Button>
        </Group>

        <Paper shadow="md" radius="md" p={0} withBorder>
          <Image
            src="https://placehold.co/800x450?text=Live+Image+Preview"
            alt="App preview"
            radius="md"
            mah={450}
            maw={800}
          />
        </Paper>
        <Badge
          leftSection={<IconBrandGolang />}
          size="md"
          variant={data?.status === 'ok' ? 'filled' : 'outline'}
        >
          {data?.status === 'ok' ? 'Server is live' : 'Server is unreachable'}
        </Badge>
      </Stack>
    </Container>
  );
}
