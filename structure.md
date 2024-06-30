# Directory structure of C:\Users\aymane\Documents\admin-ecommerce

```
├── .dirignore
├── .env
├── .eslintrc.json
├── .gitignore
├── .npmrc
├── README.md
├── actions
│   ├── get-graph-revenue.ts
│   ├── get-sales-count.ts
│   ├── get-stock-count.ts
│   └── get-total-revenue.ts
├── app
│   ├── (auth)
│   │   ├── (routes)
│   │   │   ├── sign-in
│   │   │   │   └── [[...sign-in]]
│   │   │   │       └── page.tsx
│   │   │   └── sign-up
│   │   │       └── [[...sign-up]]
│   │   │           └── page.tsx
│   │   └── layout.tsx
│   ├── (dashboard)
│   │   └── [storeId]
│   │       ├── (routes)
│   │       │   ├── billboards
│   │       │   │   ├── [billboardId]
│   │       │   │   │   ├── components
│   │       │   │   │   │   └── billboard-form.tsx
│   │       │   │   │   └── page.tsx
│   │       │   │   ├── components
│   │       │   │   │   ├── cell-action.tsx
│   │       │   │   │   ├── client.tsx
│   │       │   │   │   └── columns.tsx
│   │       │   │   └── page.tsx
│   │       │   ├── categories
│   │       │   │   ├── [categoryId]
│   │       │   │   │   ├── components
│   │       │   │   │   │   └── category-form.tsx
│   │       │   │   │   └── page.tsx
│   │       │   │   ├── components
│   │       │   │   │   ├── cell-action.tsx
│   │       │   │   │   ├── client.tsx
│   │       │   │   │   └── columns.tsx
│   │       │   │   └── page.tsx
│   │       │   ├── colors
│   │       │   │   ├── [colorId]
│   │       │   │   │   ├── components
│   │       │   │   │   │   └── size-form.tsx
│   │       │   │   │   └── page.tsx
│   │       │   │   ├── components
│   │       │   │   │   ├── cell-action.tsx
│   │       │   │   │   ├── client.tsx
│   │       │   │   │   ├── columns.tsx
│   │       │   │   │   └── delete-all-selected.tsx
│   │       │   │   └── page.tsx
│   │       │   ├── orders
│   │       │   │   ├── components
│   │       │   │   │   ├── client.tsx
│   │       │   │   │   └── columns.tsx
│   │       │   │   └── page.tsx
│   │       │   ├── page.tsx
│   │       │   ├── products
│   │       │   │   ├── [productId]
│   │       │   │   │   ├── components
│   │       │   │   │   │   └── product-form.tsx
│   │       │   │   │   └── page.tsx
│   │       │   │   ├── components
│   │       │   │   │   ├── cell-action.tsx
│   │       │   │   │   ├── client.tsx
│   │       │   │   │   └── columns.tsx
│   │       │   │   └── page.tsx
│   │       │   ├── settings
│   │       │   │   ├── components
│   │       │   │   │   └── settings-form.tsx
│   │       │   │   └── page.tsx
│   │       │   └── sizes
│   │       │       ├── [sizeId]
│   │       │       │   ├── components
│   │       │       │   │   └── size-form.tsx
│   │       │       │   └── page.tsx
│   │       │       ├── components
│   │       │       │   ├── cell-action.tsx
│   │       │       │   ├── client.tsx
│   │       │       │   ├── columns.tsx
│   │       │       │   └── delete-all-selected.tsx
│   │       │       └── page.tsx
│   │       └── layout.tsx
│   ├── (root)
│   │   ├── (routes)
│   │   │   └── page.tsx
│   │   └── layout.tsx
│   ├── api
│   │   ├── [storeId]
│   │   │   ├── billboards
│   │   │   │   ├── [billboardId]
│   │   │   │   │   └── route.ts
│   │   │   │   └── route.ts
│   │   │   ├── categories
│   │   │   │   ├── [categoryId]
│   │   │   │   │   └── route.ts
│   │   │   │   └── route.ts
│   │   │   ├── checkout
│   │   │   │   └── route.ts
│   │   │   ├── colors
│   │   │   │   ├── [colorId]
│   │   │   │   │   └── route.ts
│   │   │   │   └── route.ts
│   │   │   ├── products
│   │   │   │   ├── [productId]
│   │   │   │   │   └── route.ts
│   │   │   │   └── route.ts
│   │   │   └── sizes
│   │   │       ├── [sizeId]
│   │   │       │   └── route.ts
│   │   │       └── route.ts
│   │   ├── stores
│   │   │   ├── [storeId]
│   │   │   │   └── route.ts
│   │   │   └── route.ts
│   │   └── webhook
│   │       └── route.ts
│   ├── favicon.ico
│   ├── globals.css
│   └── layout.tsx
├── components
│   ├── main-nav.tsx
│   ├── modals
│   │   ├── alert-modal.tsx
│   │   └── store-model.tsx
│   ├── navbar.tsx
│   ├── overview.tsx
│   ├── store-switcher.tsx
│   └── ui
│       ├── Heading.tsx
│       ├── Theme-switcher.tsx
│       ├── Typography.tsx
│       ├── alert.tsx
│       ├── api-alert.tsx
│       ├── api-list.tsx
│       ├── badge.tsx
│       ├── button.tsx
│       ├── card.tsx
│       ├── checkbox.tsx
│       ├── command.tsx
│       ├── data-table.tsx
│       ├── dialog.tsx
│       ├── dropdown-menu.tsx
│       ├── form.tsx
│       ├── image-upload.tsx
│       ├── input.tsx
│       ├── label.tsx
│       ├── modal.tsx
│       ├── popover.tsx
│       ├── select.tsx
│       ├── separator.tsx
│       ├── table.tsx
│       ├── toast.tsx
│       ├── toaster.tsx
│       └── use-toast.ts
├── components.json
├── hooks
│   ├── use-origin.tsx
│   └── use-store-modal.tsx
├── lib
│   ├── prismadb.ts
│   ├── stripe.ts
│   └── utils.ts
├── middleware.ts
├── next-env.d.ts
├── next.config.js
├── package.json
├── pnpm-lock.yaml
├── postcss.config.js
├── prisma
│   └── schema.prisma
├── providers
│   ├── modal-provider.tsx
│   ├── theme-provider.tsx
│   └── toast-provider.tsx
├── public
│   ├── next.svg
│   └── vercel.svg
├── tailwind.config.js
├── tailwind.config.ts
└── tsconfig.json
```
