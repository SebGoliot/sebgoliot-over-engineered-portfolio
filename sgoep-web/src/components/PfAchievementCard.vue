<template>
  <nav class="card-list section-margin">
    <div
      class="thing-card card"
      v-for="(ach, key) in achievements.slice().reverse()"
      :key="key"
    >
      <header>
        <div>
          <span class="material-icons-round">{{ ach.icon }}</span>
          <h3 class="card-title">&nbsp;{{ ach.name }}</h3>
          <div class="availability">
            <img
              :src="require('@/assets/svg/' + ach.link_icon + '-icon.svg')"
              :alt="ach.link_icon"
            />
          </div>
        </div>
        <h4>{{ ach.subtitle }}</h4>
      </header>
      <p>
        <span v-for="(line, key) in ach.desc.split('\\n')" :key="key">
          {{ line }} <br />
        </span>
      </p>
      <div class="tech">
        <img
          v-for="(t, key) in ach.tech"
          :key="key"
          :src="require('@/assets/svg/' + t + '-icon.svg')"
          :alt="t"
        />
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: "PfAchievementCard",
  props: {
    achievements: {},
  },
};
</script>

<style scoped></style>
