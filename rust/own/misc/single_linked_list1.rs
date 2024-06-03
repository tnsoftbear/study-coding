use std::mem;

struct Node<T> {
    value: T,
    next: Option<Box<Node<T>>>
}

struct LinkedList<T> {
    head: Option<Box<Node<T>>>
}

impl<T> LinkedList<T> {
    fn new() -> Self {
        Self {
            head: None
        }
    }

    fn push(&mut self, value: T) {
        let new_node = Node { 
            value,
            next: mem::replace(&mut self.head, None)
        };
        self.head = Some(Box::new(new_node));
    }

    fn pop(&mut self) -> Option<T> {
        match mem::replace(&mut self.head, None) {
            None => None,
            Some(node_box) => {
                self.head = node_box.next;
                Some(node_box.value)
            }
        }
    }

    fn find_tail(&self) -> Option<&Node<T>> {
        if let Some(mut node_box) = self.head.as_ref() {
            while node_box.next.is_some() {
                node_box = match node_box.next.as_ref() {
                    Some(n) => n,
                    None => panic!("unexpected")
                }
            }
            Some(node_box)
        } else {
            None
        }
    }

    fn find_tail_value(&self) -> Option<&T> {
        let mut current = self.head.as_ref(); // Начинаем с головы списка
        while let Some(node) = current {
            if node.next.is_none() {
                return Some(&node.value);
            }
            current = node.next.as_ref();
        }
        None
    }

}

impl<T> Drop for LinkedList<T> {
    fn drop(&mut self) {
        let mut curr_opt = mem::replace(&mut self.head, None);
        while let Some(mut node_box) = curr_opt {
            curr_opt = mem::replace(&mut node_box.next, None);
        }
    }
}

fn main() {
    // let n1 = Node { value: 10, next: None };
    // let n2 = Node { value: 20, next: Some(Box::new(n1)) };
    // let n3 = Node { value: 30, next: Some(Box::new(n2)) };
    // let n4 = Node { value: 40, next: None };
    // let list = LinkedList { head: Some(Box::new(n3)) };
}

#[cfg(test)]
mod test {
    use super::LinkedList;

    #[test]
    fn basics() {
        let mut list = LinkedList::new();
        list.push(10);
        list.push(20);
        assert_eq!(list.pop(), Some(20));
        list.push(30);
        assert_eq!(list.find_tail().unwrap().value, 10);
        assert_eq!(list.pop(), Some(30));
        assert_eq!(list.pop(), Some(10));
        assert_eq!(list.pop(), None);
        assert!(matches!(list.find_tail(), None));
    }
}
