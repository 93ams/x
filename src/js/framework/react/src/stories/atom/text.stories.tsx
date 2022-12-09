import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import Text from "../../components/atom/text";

export default {
  title: 'Atom/Text',
  component: Text,
  argTypes: {
    message: { control: 'text' },
  },
} as ComponentMeta<typeof Text>;

const Template: ComponentStory<typeof Text> = (args) => <Text {...args} />;

export const Primary = Template.bind({});
Primary.args = {
  message: 'Bazinga!'
};
