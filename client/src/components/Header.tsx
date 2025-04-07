'use client';

import { Flex } from '@mantine/core';
import Logo from './Logo';

export default function Header() {
  return (
    <Flex justify="space-evenly" align="center">
      <Logo />
    </Flex>
  );
}
