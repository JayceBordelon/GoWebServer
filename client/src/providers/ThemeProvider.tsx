'use client';

import { MantineProvider, createTheme } from '@mantine/core';
import { useServerInsertedHTML } from 'next/navigation';
import { ServerStyles, createStylesServer } from '@mantine/next';

const stylesServer = createStylesServer();

const theme = createTheme({
  /** your custom theme config */
});

export function ThemeProvider({ children }: { children: React.ReactNode }) {
  useServerInsertedHTML(() => <ServerStyles server={stylesServer} html={'html'} />);
  return <MantineProvider theme={theme}>{children}</MantineProvider>;
}
