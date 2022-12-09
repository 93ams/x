import Text from "../components/atom/text.vue";

export default {
  title: 'Atom/Text',
  component: Text,
  argTypes: {
    message: { control: 'text' },
  },
};

const Template = (args, { argTypes }) => ({
  props: Object.keys(argTypes),
  components: { Text },
  template: '<text v-bind="$props" />',
});

export const Simple = Template.bind({});

Simple.args = {
  message: 'Bazinga'
};
