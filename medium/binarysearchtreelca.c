#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

struct Node {
	int data;
	struct Node *left, *right;
};

struct Node *newNode(int data) {
	struct Node *node = malloc(sizeof(struct Node));
	assert(node);
	node->data = data;
	node->left = node->right = NULL;

	return node;
};

void PrintTree(struct Node *root) {
	printf("\t\troot %d %p\n", root->data, root);
	printf("\tr l %d %p\t\t", root->left->data, root->left);
	printf("r r %d %p\t\t\n", root->right->data, root->right);
	printf("r l l %d %p\t", root->left->left->data, root->left->left);
	printf("r l r %d %p\t\n", root->left->right->data, root->left->right);

	return;
};

int main(int argc, char *argv[]) {
	struct Node *root = newNode(1);
	root->left = newNode(2);
	root->right = newNode(3);
	root->left->left = newNode(4);
	root->left->right = newNode(5);

	printf("\t\t *** BINARY TREE ***\n");
	PrintTree(root);

	printf("*** Recursive Traverse\n");
	printf("\n");

	return 0;
};
