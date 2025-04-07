'use client';

import { Flex, Image, Text, useMantineTheme } from '@mantine/core';

export default function Logo() {
  const theme = useMantineTheme();
  return (
    <Flex align="center" justify="space-between" py="20px" miw="145px">
      <Image src="/files.svg" alt="logo" mah="25px" />
      <Text
        component="h1"
        variant="gradient"
        gradient={theme.defaultGradient}
        tt="uppercase"
        size="xl"
        style={{ letterSpacing: '3px' }}
      >
        OPTIFILE
      </Text>
    </Flex>
  );
}
