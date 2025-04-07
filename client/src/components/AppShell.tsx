'use client';

import { AppShell as MantineAppShell, Container } from '@mantine/core';
import { ReactNode } from 'react';
import Header from './Header';

export default function AppShell({ children }: { children: ReactNode }) {
  return (
    <MantineAppShell header={{ height: 60 }} padding="md" withBorder={false}>
      <MantineAppShell.Header>
        <Header />
      </MantineAppShell.Header>

      <MantineAppShell.Main>
        <Container size="lg">{children}</Container>
      </MantineAppShell.Main>
    </MantineAppShell>
  );
}
