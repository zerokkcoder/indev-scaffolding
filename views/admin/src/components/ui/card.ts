import { cva } from 'class-variance-authority'

export const cardVariants = cva(
  'rounded-xl border bg-card text-card-foreground shadow',
  {
    variants: {
      variant: {
        default: 'shadow-sm',
        destructive:
          'bg-destructive text-destructive-foreground shadow-destructive',
      },
    },
    defaultVariants: {
      variant: 'default',
    },
  },
)

export const cardHeaderVariants = cva('flex flex-col space-y-1.5 p-6', {
  variants: {
    variant: {
      default: '',
    },
  },
  defaultVariants: {
    variant: 'default',
  },
})

export const cardTitleVariants = cva(
  'font-semibold leading-none tracking-tight',
  {
    variants: {
      variant: {
        default: '',
      },
    },
    defaultVariants: {
      variant: 'default',
    },
  },
)

export const cardDescriptionVariants = cva('text-sm text-muted-foreground', {
  variants: {
    variant: {
      default: '',
    },
  },
  defaultVariants: {
    variant: 'default',
  },
})

export const cardContentVariants = cva('p-6 pt-0', {
  variants: {
    variant: {
      default: '',
    },
  },
  defaultVariants: {
    variant: 'default',
  },
})

export const cardFooterVariants = cva('flex items-center p-6 pt-0', {
  variants: {
    variant: {
      default: '',
    },
  },
  defaultVariants: {
    variant: 'default',
  },
})
