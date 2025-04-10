<script>
    import Fa from 'svelte-fa';
    import {faRotate, faCopy} from '@fortawesome/free-solid-svg-icons';
    import {onMount} from 'svelte';
    import {toastSuccess} from '../services/Toast';
    import {GenerateUUID, GenerateULID} from '../../wailsjs/go/main/IDTab';

    let uuid = '';
    let ulid = '';
    let uuidUpper = true;
    let ulidUpper = true;

    onMount(async () => {
        GenerateUUID(uuidUpper).then((/** @type {string} */ result) => (uuid = result));
        GenerateULID(ulidUpper).then((/** @type {string} */ result) => (ulid = result));
    });

    function onGenerateUUID() {
        GenerateUUID(uuidUpper).then((/** @type {string} */ result) => (uuid = result));
    }

    function onGenerateULID() {
        GenerateULID(ulidUpper).then((/** @type {string} */ result) => (ulid = result));
    }

    function onCopy(/** @type {string} */ str) {
        navigator.clipboard.writeText(str);
        toastSuccess('Copied to clipboard!', true);
    }
</script>

<!-- svelte-ignore a11y-label-has-associated-control -->
<div class="row">
    <div class="col">
        <label class="form-label"><b>UUIDv4</b></label>
        <div class="input-group mb-3">
            <input type="text" class="form-control mono" bind:value={uuid} />
            <button class="btn btn-outline-secondary" type="button" on:click={() => onCopy(uuid)}>
                <Fa icon={faCopy} />
            </button>
        </div>
        <div class="form-check form-switch mb-3">
            <input class="form-check-input" type="checkbox" role="switch" id="uuidUpper" bind:checked={uuidUpper} />
            <label class="form-check-label" for="uuidUpper">Upper case</label>
        </div>
        <div class="d-grid">
            <button type="button" class="btn btn-primary" on:click={onGenerateUUID}>
                <Fa icon={faRotate} />
                Generate
            </button>
        </div>
    </div>
    <div class="col">
        <label class="form-label"><b>ULID</b></label>
        <div class="input-group mb-3">
            <input type="text" class="form-control mono" bind:value={ulid} />
            <button class="btn btn-outline-secondary" type="button" on:click={() => onCopy(ulid)}>
                <Fa icon={faCopy} />
            </button>
        </div>
        <div class="form-check form-switch mb-3">
            <input class="form-check-input" type="checkbox" role="switch" id="ulidUpper" bind:checked={ulidUpper} />
            <label class="form-check-label" for="ulidUpper">Upper case</label>
        </div>
        <div class="d-grid">
            <button type="button" class="btn btn-primary" on:click={onGenerateULID}>
                <Fa icon={faRotate} />
                Generate
            </button>
        </div>
    </div>
</div>
