# screen1

---
background: black
text: white
link: red+underline
vars:
    fox_is:
        set: []
---

First there was a [fox](#fox).

# fox

---
vars:
    fox_is:
        append: [pretty]
---

It was a very pretty fox.

## But also...

---
vars:
    fox_is:
        append: [smart, clever]
---

It was a very ~{smart}{clever} ~{fox.}{fox. Perhaps too clever.}

# result

The fox is {{ fox_is }}
