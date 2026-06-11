// hunk: @@ -8,7 +8,7 @@ import {createElementFromHTML, createElementFromAttrs} from '../utils/dom.ts';
import {isImageFile, isVideoFile} from '../utils.ts';
import type {DropzoneFile, DropzoneOptions} from 'dropzone/index.js';

const {i18n} = window.config;

type CustomDropzoneFile = DropzoneFile & {uuid: string};

// hunk: @@ -73,7 +73,6 @@ export async function initDropzone(dropzoneEl: HTMLElement) {
  let fileUuidDict: FileUuidDict = {}; // to record: if a comment has been saved, then the uploaded files won't be deleted from server when clicking the Remove in the dropzone
  const opts: Record<string, any> = {
    url: dropzoneEl.getAttribute('data-upload-url'),
    acceptedFiles: ['*/*', ''].includes(dropzoneEl.getAttribute('data-accepts')!) ? null : dropzoneEl.getAttribute('data-accepts'),
    addRemoveLinks: true,
    dictDefaultMessage: dropzoneEl.getAttribute('data-default-message'),
