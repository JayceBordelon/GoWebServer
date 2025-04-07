import { createTheme, rem } from '@mantine/core';

export const theme = createTheme({
  primaryColor: 'purple',

  colors: {
    pink: [
      '#fff0f5',
      '#ffd6e7',
      '#ffadcd',
      '#ff84b2',
      '#ff5b98',
      '#ff2970', // main logo pink
      '#e02664',
      '#c02158',
      '#a01c4c',
      '#801741',
    ],
    purple: [
      '#f5f4ff',
      '#ddd9ff',
      '#c3beff',
      '#a7a1f6',
      '#8a84ea',
      '#5349b7', // main logo purple
      '#473da3',
      '#3b338f',
      '#2f297a',
      '#231f66',
    ],
    blue: [
      '#e8f8ff',
      '#c0ecff',
      '#97dfff',
      '#6dd2ff',
      '#44c6ff',
      '#2eb9ff', // main logo cyan
      '#00a3ef',
      '#0097f4', // second main blue
      '#0073b3',
      '#005272',
    ],
  },

  primaryShade: { light: 5, dark: 6 },

  fontFamily: 'Inter, sans-serif',
  fontFamilyMonospace: 'Fira Code, monospace',

  headings: {
    fontFamily: 'Fredoka, sans-serif',
    fontWeight: '600',
    textWrap: 'balance',
    sizes: {
      h1: { fontSize: rem(36), lineHeight: '1.2', fontWeight: '700' },
      h2: { fontSize: rem(30), lineHeight: '1.3', fontWeight: '600' },
      h3: { fontSize: rem(24), lineHeight: '1.4', fontWeight: '600' },
      h4: { fontSize: rem(20), lineHeight: '1.4', fontWeight: '600' },
      h5: { fontSize: rem(18), lineHeight: '1.5', fontWeight: '500' },
      h6: { fontSize: rem(16), lineHeight: '1.5', fontWeight: '500' },
    },
  },

  defaultGradient: {
    from: 'pink',
    to: 'blue',
    deg: 45,
  },

  spacing: {
    xs: rem(6),
    sm: rem(12),
    md: rem(20),
    lg: rem(28),
    xl: rem(40),
  },

  fontSizes: {
    xs: rem(12),
    sm: rem(14),
    md: rem(16),
    lg: rem(18),
    xl: rem(20),
  },

  lineHeights: {
    xs: '1.4',
    sm: '1.5',
    md: '1.6',
    lg: '1.65',
    xl: '1.7',
  },

  radius: {
    xs: rem(4),
    sm: rem(8),
    md: rem(12),
    lg: rem(16),
    xl: rem(20),
  },
  defaultRadius: 'md',

  shadows: {
    xs: '0 1px 2px rgba(0, 0, 0, 0.04)',
    sm: '0 2px 4px rgba(0, 0, 0, 0.06)',
    md: '0 4px 8px rgba(0, 0, 0, 0.08)',
    lg: '0 8px 16px rgba(0, 0, 0, 0.10)',
    xl: '0 12px 24px rgba(0, 0, 0, 0.12)',
  },

  focusRing: 'auto',
  autoContrast: true,
  luminanceThreshold: 0.3,
  fontSmoothing: true,
  respectReducedMotion: true,
  cursorType: 'pointer',
});
