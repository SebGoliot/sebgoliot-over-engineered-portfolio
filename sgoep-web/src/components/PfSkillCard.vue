<template>
  <nav class="card-list section-margin">
    <div class="skill-card card" v-for="(skill, key) in skills" :key="key">
      <header>
        <img
          :src="require('@/assets/svg/'+ skill.icon +'-icon.svg')"
          :alt="skill.title"
        />
        <h3 class="card-title">&nbsp;{{ skill.title }}</h3>
      </header>
      <p>
        <span v-for="(line, key) in skill.desc.split('\\n')" :key="key">
          {{ line }} <br />
        </span>
      </p>
    </div>
  </nav>
</template>

<script>
export default {
  name: "PfSkillCard",
  props: {
    skills: {},
  },
};
</script>

<style scoped></style>

