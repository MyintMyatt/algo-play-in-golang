<script lang="ts">
  import { push } from "svelte-spa-router";
  import userProfile from "../../public/default-profile.jpeg";
  import { onMount } from "svelte";

  let currentTheme = $state<'light' | 'dark'>('dark');

  onMount(() => {
    const savedTheme = localStorage.getItem('theme') as 'light' | 'dark' | null;
    const systemPrefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
    
    if (savedTheme) {
      currentTheme = savedTheme;
    } else {
      currentTheme = systemPrefersDark ? 'dark' : 'light';
    }
    document.documentElement.setAttribute('data-theme', currentTheme);
  });

  function toggleTheme() {
    currentTheme = currentTheme === 'light' ? 'dark' : 'light';
    localStorage.setItem('theme', currentTheme);
    document.documentElement.setAttribute('data-theme', currentTheme); 
  }
</script>

<main>
  <!-- Header Section -->
  <header class="header-section">
    <h1 class="app-title primary-text">DSA Visualizations, Git, SQL Labs</h1>
    <p class="secondary-text">Learn computer science by interacting with runtime states, pointer structures, and visual data engines.</p>
  </header>

  <!-- Body Section -->
  <section class="main-body-section">
    <p class="primary-text">body</p>
  </section>

  <!-- Footer Section -->
  <footer class="footer-section">
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div class="user-profile" onclick={() => push("/profile")}>
      <!-- svelte-ignore a11y_missing_attribute -->
      <img class="profile-img" src={userProfile} />
      <p class="primary-text">nyinyimyintmyat</p>
    </div>

    <!-- Theme Switcher -->
    <button 
      type="button" 
      class="theme-button" 
      onclick={toggleTheme} 
      aria-label="Toggle light/dark theme"
    >
      <i class="theme-icon fi {currentTheme === 'light' ? 'fi-rr-sun' : 'fi-rr-moon-stars'}"></i>
    </button>

    <!-- Settings Button -->
    <button 
      type="button" 
      class="theme-button" 
      aria-label="Open settings"
      onclick={() => push("/profile")}
    >
      <i class="theme-icon fi fi-rs-settings"></i>
    </button>
  </footer>
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    width: 100%;
    padding: 30px 10px;
    box-sizing: border-box;
  }

  .app-title {
    text-align: center;
  }

  .main-body-section {
    flex-grow: 1;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  /* Footer Layout */
  .footer-section {
    width: 100%;
    padding: 0 10px;
    display: flex;
    align-items: center;
    gap: 8px;
    justify-content: flex-start;
  }

  .user-profile {
    width: fit-content;
    height: 55px;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
    padding: 8px 12px;
    border-radius: 10px;
    background-color: var(--bg-card);
    cursor: pointer;
  }

  /* Interactive Buttons replacing raw div click handlers */
  .theme-button {
    width: 55px;
    height: 55px;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0;
    border: none;
    border-radius: 10px;
    background-color: var(--bg-card);
    cursor: pointer;
    transition: background-color 0.2s ease;
  }

  .theme-button:hover {
    opacity: 0.9;
  }

  /* Flaticon Font Sizing & Centering */
  .theme-icon, i[class*="fi-tr-"] {
    font-size: 24px;
    color: var(--text-primary);
    display: inline-flex;
    align-items: center;
    justify-content: center;
    line-height: 1;
  }

  .profile-img {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    object-fit: cover;
  }
</style>