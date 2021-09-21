<template>
  <nav class="card-list section-margin">
    <div class="skill-card card" v-for="(int, key) in interests" :key="key">
      <header>
        <img
          :src="require('@/assets/svg/' + int.icon + '-icon.svg')"
          :alt="int.icon"
        />
        <h3 class="card-title">&nbsp;{{ int.name }}</h3>
      </header>
      <p>
        <span v-for="(line, key) in int.desc.split('\\n')" :key="key">
          {{ line }} <br />
        </span>
      </p>
    </div>
  </nav>
</template>

<script>
export default {
  name: "PfInterestCard",
  props: {
    interests: {},
  },
};
</script>

<style scoped></style>
