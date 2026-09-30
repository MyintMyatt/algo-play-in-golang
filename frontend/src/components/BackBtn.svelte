<script lang="ts">
    import { pop } from "svelte-spa-router";
</script>

<section class="back-btn-section">
    <!-- svelte-ignore a11y_consider_explicit_label -->
    <button class="back-btn" onclick={() => pop()}>
        <i class="back-icon fi fi-br-angle-double-small-left"></i>
    </button>
</section>

<style>
    .back-btn-section {
        display: flex;
        width: 100%;
        margin: 0;
        margin-left: 40px;
        position: fixed;
        top: 50px;
        justify-content: start;
        align-items: center;
    }

    .back-btn {
        font-size: 25px;
        padding: 10px;
        background-color: var(--accent-primary);
        border-radius: 10px;
        border: none;
        transition: all 0.2s ease;
    }

    .back-btn:hover {
        transform: translate(2px, 2px);
        box-shadow: 2px 2px 0px var(--text-secondary);
    }
    .back-btn:active {
        transform: translate(4px, 4px);
        box-shadow: 0px 0px 0px #000000;
    }

    .back-icon {
        color: var(--text-primary);
    }

    .back-icon:hover {
        color: var(--bg-canvas);
        transition: all 0.3s ease-in-out;
    }
</style>