/** @type {import("@builder.io/mitosis/dist/src/generators/qwik/component-generator").ToQwikOptions} */
const quikOpts = {
    typescript: true,
}
/** @type {import("@builder.io/mitosis").ToReactOptions} */
const reactOpts = {
    typescript: true,
}
/** @type {import("@builder.io/mitosis").ToVueOptions} */
const vueOpts = {
    typescript: true
}

/** @type {import("@builder.io/mitosis").MitosisConfig} */
module.exports = {
    files: ['./atom/**'],
    targets: ['vue3', 'qwik', 'react'],
    options: {
        react: reactOpts,
        qwik: quikOpts,
        vue3: vueOpts,
    },
    getTargetPath: (target) => {
        return `../../framework/${target.target}/src/components`
    }
}