<template>
  <li>
    {{ todo.text }}
    <button @click="$emit('remove', todo.id)">X</button>
  </li>
</template>

<script>
module.exports = {
  props: {
    todo: {
      type: Object,
      required: true,
    },
  },
};
</script>