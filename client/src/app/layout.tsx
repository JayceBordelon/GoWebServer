import '@mantine/core/styles.css';
import type { Metadata } from 'next';
import { ThemeProvider } from '@providers/ThemeProvider';

export const metadata: Metadata = {
  title: 'Optifile',
  description: 'Store, serve, get, transform, and share files with ease.',
  icons: ['/file.svg'],
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>
        <ThemeProvider>{children}</ThemeProvider>
      </body>
    </html>
  );
}
