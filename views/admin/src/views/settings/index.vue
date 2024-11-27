<script setup lang="ts">
// Basic settings page
import { ref } from 'vue'

const settings = ref({
  theme: 'light',
  notifications: true
})

const updateSettings = () => {
  // Placeholder for settings update logic
  console.log('Settings updated:', settings.value)
}
</script>

<template>
  <div class="settings-page">
    <h1>Application Settings</h1>
    <div class="settings-form">
      <label>
        Theme:
        <select v-model="settings.theme">
          <option value="light">Light</option>
          <option value="dark">Dark</option>
        </select>
      </label>
      <label>
        Notifications:
        <input type="checkbox" v-model="settings.notifications" />
      </label>
      <button @click="updateSettings">Save Settings</button>
    </div>
  </div>
</template>

<style scoped>
.settings-page {
  padding: 20px;
}
.settings-form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
</style>
