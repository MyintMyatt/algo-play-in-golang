<script lang="ts">
    import { onMount } from "svelte";
    import userProfile from "../../public/default-profile.jpeg";
    import BackBtn from "../components/BackBtn.svelte";
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
    <BackBtn />

    <section class="heading-section">
       <div class="profile-heading">
            <!-- svelte-ignore a11y_missing_attribute -->
            <img class="profile-img" src={userProfile} />
            <h3>nyinyimyintmyat</h3>
       </div> 
    </section>

    <!-- about -->
    <section class="about-section">
        <div class="about-user">
            <p class="sub-title">About</p>
            <hr class="break-line" />
            <div class="data-row">
                <p class="app-label">Full Name:</p>
                <p class="about-data-value">Nyi Nyi Myint Myat</p>
            </div>

            <div class="data-row">
                <p class="app-label">Nick Name:</p>
                <p class="about-data-value">orion</p>
            </div>

            <div class="data-row">
                <p class="app-label">Tech Stack:</p>
                <p class="about-data-value">java, golang,...</p>
            </div>
        </div>
    </section>

    <!-- setting -->
    <section class="setting-section">
       <div class="setting-content">
            <p class="sub-title">Setting</p>
            <hr class="break-line" />
            <div class="setting-row">
                <p class="app-label">Theme</p>
                <!-- svelte-ignore a11y_consider_explicit_label -->
                <button class="theme-icon" onclick={toggleTheme}>
                    <i class="theme-icon fi {currentTheme === 'light' ? 'fi-rr-sun' : 'fi-rr-moon-stars'}"></i>
                </button>
            </div>

            <div class="setting-row">
                <p class="app-label">Sound Effect</p>
                <p class="about-data-value">orion</p>
            </div>

            <div class="setting-row">
                <p class="app-label">Tech Stack:</p>
                <p class="about-data-value">java, golang,...</p>
            </div>
        </div> 
    </section>

    <section class="setting-section">

    </section>
</main>

<style>
  * {
    margin: 0;
  }

  main {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    width: 100%;
    overflow-y: scroll;
    gap: 35px;
    padding: 30px 10px;
    box-sizing: border-box;
  }

  .heading-section, .about-section, .setting-section {
    background-color: transparent;
    width: 100%;
    height: fit-content;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .about-section {
    margin-top: 30px;
  }

  .profile-heading, .about-user, .setting-content {
    height: 100%;
    width: 500px;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .profile-img {
    width: 180px;
    height: 180px;
    border-radius: 150px;
    border: 2px dotted var(--border-strong);
    padding: 1px;
    margin-bottom: 20px;
  }

  .about-user, .setting-content{
    background-color: var(--bg-surface);
    box-shadow: 1px 2px var(--accent-primary);
    border: 1px solid var(--border-subtle);
    align-items: start;
    justify-content: start;
    gap: 10px;
    margin: 0;
    padding: 25px;
  }

  .break-line {
    color: var(--border-subtle);
    background-color: red;
    width: 100%;
  }

  .data-row, .setting-row{
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: var(--bg-card);
    padding: 10px;
    border-radius: 10px;
    height: 50px;
  }

  /* .sub-title {
    width: 100px;
  } */

  .app-label {
    color: var(--text-secondary);
    margin: 2px;
    width: 120px;
  }

  .about-data-value {
    display: flex;
    font-weight: bold;
    justify-content: end;
    flex: 1;
  }

  /* setting */
  .setting-section {
    height: fit-content;
  }
  
  .setting-row {
    background-color: var(--bg-card);
    padding: 10px 10px;
    border-radius: 10px;
    height: 50px;
  }

  .theme-icon, i[class*="fi-tr-"] {
    font-size: 24px;
    color: var(--text-primary);
    background-color: transparent;
    border: 0;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    line-height: 1;
  }

  .theme-icon:hover {
    font-size: 22px;
  }

</style>