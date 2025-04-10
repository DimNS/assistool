<script>
    import Fa from 'svelte-fa';
    import {faPlay, faEraser, faCopy} from '@fortawesome/free-solid-svg-icons';
    import {onMount} from 'svelte';
    import {Modal as BSModal} from 'bootstrap/dist/js/bootstrap.esm';
    import Modal from '../components/Modal.svelte';
    import {toastSuccess} from '../services/Toast';
    import {Beautify} from '../../wailsjs/go/main/JSONTab';

    const clearModalID = 'json-clear-modal';

    let /** @type {BSModal} */ clearModal;

    let json = '';
    let ntconvert = true;

    onMount(async () => {
        clearModal = new BSModal(document.getElementById(clearModalID));
    });

    function onBeautify() {
        Beautify(json, ntconvert).then((result) => (json = result));
    }

    function onCopy() {
        navigator.clipboard.writeText(json);
        toastSuccess('Copied to clipboard!', true);
    }

    function onClear() {
        json = '';
        clearModal.hide();
    }
</script>

<button type="button" class="btn btn-outline-secondary mt-3 json-copy" on:click={onCopy}>
    <Fa icon={faCopy} />
</button>
<textarea
    class="form-control json-textarea mono"
    placeholder="Place your json string here..."
    bind:value={json}
></textarea>
<div class="form-check form-switch mt-3">
    <input class="form-check-input" type="checkbox" role="switch" id="ntconvert" bind:checked={ntconvert} />
    <label class="form-check-label" for="ntconvert">Convert \n and \t</label>
</div>
<div class="row mt-3">
    <div class="col-8">
        <div class="d-grid">
            <button type="button" class="btn btn-primary" on:click={onBeautify}>
                <Fa icon={faPlay} />
                Beautify
            </button>
        </div>
    </div>
    <div class="col-4">
        <div class="d-grid">
            <button type="button" class="btn btn-light" on:click={() => clearModal.show()}>
                <Fa icon={faEraser} />
                Clear
            </button>
        </div>
    </div>
</div>

<Modal id={clearModalID} size="md" title="Clearing the field">
    <div slot="body">Are you sure you want to clear the field?</div>
    <div slot="footer">
        <button type="button" on:click={onClear} class="btn btn-success">
            <Fa icon={faEraser} />
            Yes
        </button>
        &nbsp;
        <button type="button" data-bs-dismiss="modal" class="btn btn-light">No</button>
    </div>
</Modal>

<style>
    .json-copy {
        position: absolute;
        right: 0;
        margin-right: 34px;
    }

    .json-textarea {
        height: calc(100vh - 184px);
        min-height: calc(100vh - 184px);
        max-height: calc(100vh - 184px);
    }
</style>
