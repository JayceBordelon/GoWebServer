import { Button, Flex } from '@mantine/core';
import { IconLogin2 } from '@tabler/icons-react';

export default function NavLinks() {
  return (
    <Flex direction="row-reverse" gap="md" align="center">
      <Button leftSection={<IconLogin2 />} variant="outline">
        Log In
      </Button>
    </Flex>
  );
}
