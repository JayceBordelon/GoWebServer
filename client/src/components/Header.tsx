'use client';

import { Flex } from '@mantine/core';
import Logo from './Logo';
import NavLinks from './NavLinks';

export default function Header() {
  return (
    <Flex justify="space-evenly" align="center" w="100%">
      <Logo />
      <NavLinks />
    </Flex>
  );
}
