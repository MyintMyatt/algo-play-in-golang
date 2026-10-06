<script lang="ts">
    interface ToggleProps {
        checked?: boolean;
        onToggle?: (checked: boolean) => void;
    }

    let { checked = false, onToggle }: ToggleProps = $props();

    function toggle() {
        checked = !checked;
        if (onToggle) {
            onToggle(checked);
        }
    }
</script>

<label class="switch">
    <!-- svelte-ignore event_directive_deprecated -->
    <input type="checkbox" {checked} on:change={toggle} />
    <span class="slider round"></span>
</label>

<style>
    .switch {
        position: relative;
        display: inline-block;
        width: 60px;
        height: 30px;
    }

    .switch input {
        opacity: 0;
        width: 0;
        height: 0;
    }

    .slider {
        position: absolute;
        cursor: pointer;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: var(--border-strong);
        -webkit-transition: 0.4s;
        transition: 0.4s;
    }

    .slider:before {
        position: absolute;
        content: "";
        height: 22px;
        width: 22px;
        left: 4px;
        bottom: 4px;
        background-color: var(--accent-primary);
        -webkit-transition: 0.4s;
        transition: 0.4s;
    }

    input:checked + .slider {
        background-color: var(--accent-primary);
    }

    input:focus + .slider {
        box-shadow: 0 0 1px #2196f3;
    }
    input:checked + .slider:before {
        -webkit-transform: translateX(26px);
        -ms-transform: translateX(26px);
        background-color: white;
        transform: translateX(26px);
    }

    /* Rounded sliders */
    .slider.round {
        border-radius: 34px;
    }

    .slider.round:before {
        border-radius: 50%;
    }
</style>
