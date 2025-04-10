<script>
    import {SvelteToast} from '@zerodevx/svelte-toast';
    import JsonTab from './tabs/Json.svelte';
    import GoStruct from './tabs/GoStruct.svelte';
    import IdTab from './tabs/Id.svelte';
    import PercentTab from './tabs/Percent.svelte';

    const svelteToastOptions = {
        duration: 5000,
        pausable: true,
    };
</script>

<main class="p-3">
    <ul class="nav nav-pills" id="main-menu" role="tablist">
        <li class="nav-item" role="presentation">
            <button
                class="nav-link active"
                id="json-tab"
                data-bs-toggle="tab"
                data-bs-target="#json-tab-pane"
                type="button"
                role="tab"
            >
                JSON
            </button>
        </li>
        <li class="nav-item" role="presentation">
            <button
                class="nav-link"
                id="gostruct-tab"
                data-bs-toggle="tab"
                data-bs-target="#gostruct-tab-pane"
                type="button"
                role="tab"
            >
                GoStruct
            </button>
        </li>
        <li class="nav-item" role="presentation">
            <button
                class="nav-link"
                id="uuidulid-tab"
                data-bs-toggle="tab"
                data-bs-target="#uuidulid-tab-pane"
                type="button"
                role="tab"
            >
                UUID/ULID
            </button>
        </li>
        <li class="nav-item" role="presentation">
            <button
                class="nav-link"
                id="percent-tab"
                data-bs-toggle="tab"
                data-bs-target="#percent-tab-pane"
                type="button"
                role="tab"
            >
                PERCENT
            </button>
        </li>
    </ul>
    <div class="tab-content mt-3" id="main-menu">
        <div class="tab-pane fade show active" id="json-tab-pane" role="tabpanel" tabindex="0">
            <JsonTab />
        </div>
        <div class="tab-pane fade" id="gostruct-tab-pane" role="tabpanel" tabindex="0">
            <GoStruct />
        </div>
        <div class="tab-pane fade" id="uuidulid-tab-pane" role="tabpanel" tabindex="0">
            <IdTab />
        </div>
        <div class="tab-pane fade" id="percent-tab-pane" role="tabpanel" tabindex="0">
            <PercentTab />
        </div>
    </div>
</main>

<SvelteToast options={svelteToastOptions} />
