import '@mantine/core/styles.css';
import '../global.css';

import React, { ReactNode } from 'react';
import { ColorSchemeScript, mantineHtmlProps, MantineProvider } from '@mantine/core';
import AppShell from '@components/AppShell';
import { theme } from '@lib/theme';

export const metadata = {
  title: 'Optifile',
  description: 'Serve, share and transform your files with ease.',
  icons: {
    icon: '/files.svg',
    shortcut: '/files.svg',
    apple: '/files.svg',
  },
};

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en" {...mantineHtmlProps}>
      <head>
        <ColorSchemeScript />
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width, user-scalable=no"
        />
      </head>
      <body>
        <MantineProvider theme={theme}>
          <AppShell>{children}</AppShell>
        </MantineProvider>
      </body>
    </html>
  );
}
