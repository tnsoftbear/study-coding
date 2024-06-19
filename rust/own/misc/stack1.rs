use std::rc::Rc;

type Link<T> = Option<Rc<Node<T>>>;

struct Node<T> {
    elem: T,
    next: Link<T>
}

struct List<T> {
    head: Link<T>
}

impl<T> List<T> {
    pub fn new() -> Self {
        List { head: None }
    }

    pub fn prepend(&mut self, elem: T) {
        let new_node = Node {
            elem,
            next: self.head.take()
        };
        self.head = Some(Rc::new(new_node));
    }
}

fn main() {}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn basics() {
        let mut list: List<i32> = List::new();
        list.prepend(10);
    }
}
